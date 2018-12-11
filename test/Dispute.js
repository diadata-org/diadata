const util = require('util');

const encodeCall = require('zos-lib/lib/helpers/encodeCall').default;
const Dispute = artifacts.require('Dispute');
const DIAToken = artifacts.require('DIAToken');
const maxTimeTol = 15
// prevent tight block times
const disputeLength = 2 * 7 * 24 * 60 * 60 + maxTimeTol;
const voteCost = 10;


advanceTimeAndBlock = async (time) => {
    let block = await web3.eth.getBlock("latest")
    await advanceTime(time);
    await advanceBlock();
    let blockN = await web3.eth.getBlock("latest")
    assert.notEqual(block.hash, blockN.hash, "Block not mined!");
    assert.equal(Math.abs(blockN.timestamp - time - block.timestamp) < maxTimeTol, true, "Block timestamp not increased");
    return Promise.resolve(web3.eth.getBlock('latest'));
}

advanceTime = (time) => {
    return new Promise((resolve, reject) => {
        web3.currentProvider.sendAsync({
            jsonrpc: "2.0",
            method: "evm_increaseTime",
            params: [time],
            id: new Date().getTime()
        }, (err, result) => {
            if (err) { return reject(err); }
            return resolve(result);
        });
    });
}

advanceBlock = () => {
    return new Promise((resolve, reject) => {
        web3.currentProvider.sendAsync({
            jsonrpc: "2.0",
            method: "evm_mine",
            id: new Date().getTime()
        }, (err, result) => {
            if (err) { return reject(err); }
            const newBlockHash = web3.eth.getBlock('latest').hash;

            return resolve(newBlockHash)
        });
    });
}

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

        await advanceTimeAndBlock(disputeLength)
        await dispute.triggerDecision(1)

        let reward = Math.floor((10 - winners) / winners * voteCost);
        vote = true;


        let balance = await dispute.checkRewardsBalance.call({ from: owner });
        assert.equal(balance.valueOf(), voteCost + reward, "Owner balance mismatch");

        for (var i = 1; i < 10; i++) {
            let balance = await dispute.checkRewardsBalance.call({ from: accounts[i] });
            assert.equal(balance.valueOf(), (vote ? voteCost + reward : 0), "account " + i.toString() + " balance mismatch");
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

        await advanceTimeAndBlock(disputeLength);

        await dispute.triggerDecision(1)

        let reward = (10 - winners) / winners * voteCost
        vote = false;

        let balance = await dispute.checkRewardsBalance.call({ from: owner });
        assert.equal(balance.valueOf(), 0, "Owner balance mismatch");

        balance = await dispute.checkRewardsBalance.call({ from: holder });
        assert.equal(balance.valueOf(), voteCost + reward, "holder balance mismatch");


        for (var i = 2; i < 10; i++) {
            balance = await dispute.checkRewardsBalance.call({ from: accounts[i] });
            assert.equal(balance.valueOf(), (vote ? 0 : voteCost + reward), "account " + i.toString() + " balance mismatch");

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
        await advanceTimeAndBlock(disputeLength)
        await dispute.triggerDecision(1, { from: notHolder })
        let ongoing = await dispute.isDisputeOpen.call(1);

        assert.equal(ongoing.valueOf(), false, "Dispute was not closed");
    });

    it("Event is emitted when dispute is closed", async function () {
        await dispute.openDispute(1);
        await advanceTimeAndBlock(disputeLength)

        tx = await dispute.triggerDecision(1, { from: notHolder })

        assert.equal(tx.logs[0].event, 'DisputeClosed');
        assert.equal(tx.logs[0].args._id.valueOf(), 1);
        assert.equal(tx.logs[0].args._result.valueOf(), true);
    });

    it("Event is emitted when dispute is created", async function () {
        let block = await web3.eth.getBlock("latest")
        let tx = await dispute.openDispute(1);

        assert.equal(tx.logs[0].event, 'DisputeOpen');
        assert.equal(tx.logs[0].args._id.valueOf(), 1);
        assert.equal(Math.abs(tx.logs[0].args._deadline.valueOf() - block.timestamp - disputeLength + maxTimeTol) < maxTimeTol, true);
    })

    it("Dispute can be opened after is closed", async function () {
        await dispute.openDispute(1);
        await advanceTimeAndBlock(disputeLength)
        await dispute.triggerDecision(1, { from: notHolder })
        await dispute.openDispute(1);
        let ongoing = await dispute.isDisputeOpen.call(1);

        assert.equal(ongoing.valueOf(), true, "Dispute was not re-opened");
    });

    it("Nobody should be able to trigger decision before deadline", async function () {
        await dispute.openDispute(1);
        await advanceTimeAndBlock(disputeLength - 100)
        await web3.eth.getBlock("latest")

        await assertRevert(dispute.triggerDecision(1));
    });

    it("Participating in multiple votes should grant multiple rewards", async function () {
        let max = 5;
        for (let i = 0; i < max; i++) {
            await dispute.openDispute(i);
        }

        await advanceTimeAndBlock(disputeLength + 10 * max)

        for (let i = 0; i < max; i++) {
            await dispute.triggerDecision(i);
        }

        balance = await dispute.checkRewardsBalance.call()
        await assert.equal(balance.valueOf(), max * voteCost, "Balance mismatch");
    });

    it("If reward balance is 0 should not be possible to withdraw", async function () {
        await assertRevert(dispute.claimRewards());
    });

    it("If reward balance is not 0 should be possible to withdraw", async function () {
        await dispute.openDispute(1);
        await advanceTimeAndBlock(disputeLength)
        await dispute.triggerDecision(1)
        await dispute.claimRewards()
        balance = await dia.balanceOf(owner)

        assert.equal(balance.valueOf(), initialBalance, "Dia holder vote not processed");
    });
});
