const util = require('util');

const encodeCall = require('zos-lib/lib/helpers/encodeCall').default;
const Dispute = artifacts.require('Dispute');
const DIAToken = artifacts.require('DIAToken');
const disputeLength = 10;//2 * 7 * 24 * 60 * 60 / 15;
const voteCost = 10;

const waitNBlocks = async n => {
    const sendAsync = util.promisify(web3.currentProvider.sendAsync);
    await Promise.all(
        [...Array(n).keys()].map(i =>
            sendAsync({
                jsonrpc: '2.0',
                method: 'evm_mine',
                id: i
            })
        )
    );
};

async function assertRevert(promise) {
    try {
        await promise;
        assert.fail('Expected revert not received');
    } catch (error) {
        const revertFound = error.message.search('revert') >= 0;
        assert(revertFound, `Expected "revert", got ${error} instead`);
    }
};

contract('Dispute', function (accounts) {
    let owner = accounts[0];
    let holder = accounts[1];
    let notHolder = accounts[2];
    let dispute;
    let initialBalance = 1e6;


    beforeEach('setup contract before each test', async function () {
        dispute = await Dispute.new({ from: owner });
        dia = await DIAToken.new({ from: owner });

        var callData = encodeCall('initialize', ['address'], [owner]);
        await dia.sendTransaction({ data: callData, from: owner });
        callData = encodeCall('initialize', ['address'], [dia.address]);
        await dispute.sendTransaction({ data: callData, from: owner });

        await dia.mint(owner, initialBalance);
        await dia.mint(holder, initialBalance);
        await dia.increaseAllowance(dispute.address, initialBalance, { from: owner });
        await dia.increaseAllowance(dispute.address, initialBalance, { from: holder });
    });

    it("DIA token holder should be able to open a dispute", async function () {
        await dispute.openDispute(1, { from: holder });
        let ongoing = await dispute.isDisputeOpen.call(1);
        assert.equal(ongoing.valueOf(), true, "Dispute was not initiated");
    });

    it("Users without DIA token should not be able to open a dispute", async function () {
        await assertRevert(dispute.openDispute(1, { from: notHolder }));
    });

    it("Dispute should not be able to be initiated if already in progress", async function () {
        await dispute.openDispute(1);

        await assertRevert(dispute.openDispute(1));
    });

    it("DIA token holder should be able to vote on a dispute", async function () {
        await dispute.openDispute(1);
        await dispute.vote(1, true, { from: holder })
        let vote = await dispute.didCastVote.call(1, { from: holder })

        assert.equal(vote.valueOf(), true, "Dia holder vote not processed");
    });

    it("DIA token holder should not be able to vote twice on a dispute", async function () {
        await dispute.openDispute(1);
        await dispute.vote(1, true, { from: holder })

        await assertRevert(dispute.vote(1, true, { from: holder }))
    });

    it("Users without DIA token should not be able to vote on a dispute", async function () {
        await dispute.openDispute(1);

        await assertRevert(dispute.vote(1, false, { from: notHolder }))
    });

    it("DIA balance should decrease after a dispute", async function () {
        await dispute.openDispute(1);
        dispute.vote(1, false, { from: holder });
        balanceOwner = await dia.balanceOf.call(owner);
        balanceHolder = await dia.balanceOf.call(holder);
        assert.equal(balanceOwner.valueOf(), initialBalance - voteCost, "Owner balance mismatch");
        assert.equal(balanceHolder.valueOf(), initialBalance - voteCost, "Holder balance mismatch");
    });

    it("DIA winners (drop) should be paid out", async function () {
        await dispute.openDispute(1);
        var vote = true;
        var winners = 2;
        dispute.vote(1, vote, { from: holder });
        vote = !vote;

        for (var i = 2; i < 10; i++) {
            await dia.mint(accounts[i], initialBalance);
            await dia.increaseAllowance(dispute.address, initialBalance, { from: accounts[i] });
            await dispute.vote(1, vote, { from: accounts[i] });
            if (vote) winners++;
            vote = !vote;
        }

        waitNBlocks(disputeLength)
        await dispute.triggerDecision(1)

        let reward = Math.floor((10 - winners) / winners * voteCost);
        vote = true;

        let balance = await dia.balanceOf.call(owner);
        assert.equal(balance.valueOf(), initialBalance + reward, "Owner balance mismatch");

        for (var i = 1; i < 10; i++) {
            balance = await dia.balanceOf.call(accounts[i]);
            assert.equal(balance.valueOf(), initialBalance + (vote ? reward : -voteCost), "account " + i.toString() + " balance mismatch");
            vote = !vote;
        }
    });

    it("DIA winners (keep) should be paid out", async function () {
        await dispute.openDispute(1);
        var vote = false;
        var winners = 1;
        dispute.vote(1, vote, { from: holder });

        for (var i = 2; i < 10; i++) {
            await dia.mint(accounts[i], initialBalance);
            await dia.increaseAllowance(dispute.address, initialBalance, { from: accounts[i] });
            await dispute.vote(1, vote, { from: accounts[i] });
            if (!vote) winners++;
            vote = !vote;
        }

        waitNBlocks(disputeLength)
        await dispute.triggerDecision(1)

        let reward = (10 - winners) / winners * voteCost
        vote = false;

        let balance = await dia.balanceOf.call(owner);
        assert.equal(balance.valueOf(), initialBalance - voteCost, "Owner balance mismatch");

        balance = await dia.balanceOf.call(holder);
        assert.equal(balance.valueOf(), initialBalance + reward, "holder balance mismatch");


        for (var i = 2; i < 10; i++) {
            balance = await dia.balanceOf.call(accounts[i]);
            assert.equal(balance.valueOf(), initialBalance + (vote ? -voteCost : reward), "account " + i.toString() + " balance mismatch");
            vote = !vote;
        }
    });

    it("Users without enough DIA token balance should not be able to vote on a dispute", async function () {
        let initialBalance = voteCost - 1;
        await dia.mint(notHolder, initialBalance);
        await dia.increaseAllowance(dispute.address, initialBalance, { from: notHolder });
        await dispute.openDispute(1);

        await assertRevert(dispute.vote(1, true, { from: notHolder }))
    });

    it("Anyone should be able to trigger decision after deadline", async function () {
        await dispute.openDispute(1);
        waitNBlocks(disputeLength);
        await dispute.triggerDecision(1, { from: notHolder })
        let ongoing = await dispute.isDisputeOpen.call(1);

        assert.equal(ongoing.valueOf(), false, "Dispute was not closed");
    });

    it("Event is emitted when dispute is closed", async function () {
        await dispute.openDispute(1);
        waitNBlocks(disputeLength);
        tx = await dispute.triggerDecision(1, { from: notHolder })

        assert.equal(tx.logs[0].event, 'DisputeClosed');
        assert.equal(tx.logs[0].args._id.valueOf(), 1);
        assert.equal(tx.logs[0].args._result.valueOf(), true);
    });

    it("Event is emitted when dispute is created", async function () {
        let tx = await dispute.openDispute(1);

        assert.equal(tx.logs[0].event, 'DisputeOpen');
        assert.equal(tx.logs[0].args._id.valueOf(), 1);
        assert.equal(tx.logs[0].args._deadline.valueOf(), tx.logs[0].blockNumber + disputeLength);
    })

    it("Dispute can be opened after is closed", async function () {
        await dispute.openDispute(1);
        waitNBlocks(disputeLength);
        await dispute.triggerDecision(1, { from: notHolder })
        await dispute.openDispute(1);
        let ongoing = await dispute.isDisputeOpen.call(1);

        assert.equal(ongoing.valueOf(), true, "Dispute was not re-opened");
    });

    it("Nobody should be able to trigger decision before deadline", async function () {
        await dispute.openDispute(1);
        waitNBlocks(disputeLength - 1);

        await assertRevert(dispute.triggerDecision(1));
    });
});
