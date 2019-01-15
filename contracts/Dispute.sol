pragma solidity 0.4.24;

import "./DIAToken.sol";
import "zos-lib/contracts/Initializable.sol";
import "openzeppelin-eth/contracts/math/SafeMath.sol";

contract Dispute is Initializable {
	using SafeMath for uint;
	// DIA token
	DIAToken private dia_;
	// Event to emit when a dispute is open
	event DisputeOpen(uint256 _id, uint _deadline);
	// Event to emit when a dispute is finalized
	event DisputeClosed(uint256 _id, bool _result);
	// How many blocks should we wait before the dispute can be closed
	uint private DISPUTE_LENGTH;
	// How many token a user should stake on each vote
	uint private VOTE_COST;
	// Disputes
	mapping (uint256=>Disputes) private disputes_;
	// Rewards that each voter can claim
	mapping (address=>uint256) public rewards_;

	struct Disputes{
		// Block number to finalize dispute
		uint deadline;
		// Array of voters
		Voter[] voters;
		// Voters index in the voters array
		mapping(address=>uint) votersIndex;
	}
	// Voters
	struct Voter {
		// store voter address. Required for payout
		address id;
		// Vote. true:keep, false:drop
		bool vote;
	}

	/**
	* @dev Acts as constructor for upgradeable contracts
	* @param _dia Address of DIA token contract.
	*/
	function initialize(DIAToken _dia) public initializer() {
		// ~2 weeks. weeks x days x hours x minute x seconds
		DISPUTE_LENGTH = 2*7*24*60*60;
		VOTE_COST = 10;
		dia_ = _dia;
	}

	/**
	* @dev Cast vote.
	* @param _id Data source identifier.
	* @param _vote true for drop and false to keep.
	*/
	function vote(uint256 _id, bool _vote) public{
		// check only new voters
		require (disputes_[_id].votersIndex[msg.sender] == 0, "Address already voted");
		require (disputes_[_id].deadline > 0, "Dispute not available");
		dia_.transferFrom(msg.sender, this, VOTE_COST);
		disputes_[_id].voters.push(Voter(msg.sender, _vote));
		disputes_[_id].votersIndex[msg.sender] = disputes_[_id].voters.length;
	}

	/**
	* @dev Start a dispute.
	* @param _id data source identifier.
	*/
	function openDispute(uint256 _id) external {
		require(disputes_[_id].deadline == 0, "Dispute already ongoing");
		disputes_[_id].deadline = now+DISPUTE_LENGTH;
		emit DisputeOpen(_id, disputes_[_id].deadline);
		vote(_id, true);
	}

	/**
	* @dev Once the deadline is reached this function should be called to get decision.
	* @param _id data source id.
	*/
	function triggerDecision(uint256 _id) external{
		// Maybe we can get rid of a require
		require(disputes_[_id].deadline > 0, "Dispute not available");
		require(now > disputes_[_id].deadline, "Dispute deadline not reached");
		// prevent method to be called again before its done
		disputes_[_id].deadline = 0;
		uint256 dropVotes = 0;
		uint256 keepVotes = 0;
		uint totalVoters = disputes_[_id].voters.length;
		for (uint i = 0; i<totalVoters; i++){
			if (disputes_[_id].voters[i].vote)
				dropVotes++;
			else
				keepVotes++;
		}
		bool drop = (dropVotes>keepVotes);
		uint payment;
		// use safe math to compute payment
		if (drop)
			payment = ((totalVoters).mul(VOTE_COST)).div(dropVotes);
		else
			payment = ((totalVoters).mul(VOTE_COST)).div(keepVotes);
		for (i = 0; i < totalVoters; i++){
			if (disputes_[_id].voters[i].vote == drop){
				rewards_[disputes_[_id].voters[i].id] += payment;
			}
			delete disputes_[_id].votersIndex[disputes_[_id].voters[i].id];
		}
		delete disputes_[_id];
		emit DisputeClosed(_id, drop);
	}

	/**
	* @dev Claim rewards
	*/
	function claimRewards() external {
		require(rewards_[msg.sender] > 0, "No balance to withdraw");
		dia_.transfer(msg.sender, rewards_[msg.sender]);
		rewards_[msg.sender] = 0;
	}

	/**
	* @dev Check rewards balance for account calling the method
	*/
	function checkRewardsBalance() external view returns (uint256) {
		return rewards_[msg.sender];
	}

	/**
	* @dev get dispute status.
	* @param _id data source id.
	*/
	function isDisputeOpen(uint256 _id) external view returns (bool) {
		return (disputes_[_id].deadline>0);
	}

	/**
	* @dev check if address voted already.
	* @param _id data source identifier.
	*/
	function didCastVote(uint256 _id) external view returns (bool){
		return (disputes_[_id].votersIndex[msg.sender]>0);
	}
}
