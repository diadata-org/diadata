pragma solidity ^0.4.24;

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
	uint private voteLength_;
	// How many token a user should stake on each vote
	uint private voteCost_;
	// Disputes
	mapping (uint256=>Disputes) private disputes_;
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
	* @dev Allows the current owner to transfer control of the contract to a newOwner.
	* @param _dia Address of DIA token contract.
	*/
	function initialize(DIAToken _dia) public initializer() {
		dia_ = _dia;
		voteLength_ = 10;
		voteCost_ = 10;
	}

	/**
	* @dev Start a dispute.
	* @param _id data source identifier.
	*/
	function openDispute(uint256 _id) public {
		require(disputes_[_id].deadline == 0, "Dispute already ongoing");
		disputes_[_id].deadline = block.number+voteLength_;
		emit DisputeOpen(_id, disputes_[_id].deadline);
		vote(_id, true);
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
		dia_.transferFrom(msg.sender, this, voteCost_);
		disputes_[_id].voters.push(Voter(msg.sender, _vote));
		// Safe math may be required for many many voters
		disputes_[_id].votersIndex[msg.sender] = disputes_[_id].voters.length;
	}

	/**
	* @dev Once the deadline is reached this function should be called to get decision.
	* @param _id data source id.
	*/
	function triggerDecision(uint256 _id) public{
		// Maybe we can get rid of a require
		// we may need to prevent trigger to be called multiple times for a dispute
		// we can delete dispute here
		// Person opening the dispute should vote?
		// Dispute dispute = disputes_[_id];
		require(disputes_[_id].deadline > 0, "Dispute not available");
		require(block.number > disputes_[_id].deadline, "Dispute deadline not reached");
		uint256 dropVotes = 0;
		uint256 keepVotes = 0;
		// if total voters did not overflow then votes neither
		uint totalVoters = disputes_[_id].voters.length;
		for (uint i = 0; i<totalVoters; i++){
			if (disputes_[_id].voters[i].vote)
				dropVotes++;
			else
				keepVotes++;
		}
		bool drop = (dropVotes>keepVotes);
		uint payment;
		if (drop)
			// Protect the payment
			payment = ((totalVoters).mul(voteCost_)).div(dropVotes);
		else
			payment = ((totalVoters).mul(voteCost_)).div(keepVotes);
		// Maybe we can move this to claim reward? so the person closing the vote does not pay transfer fees
		// Other options is to give vote remanent to person closing the vote
		// uint winners=0;
		for (i = 0; i < totalVoters; i++){
			if (disputes_[_id].voters[i].vote == drop){
				dia_.transfer(disputes_[_id].voters[i].id, payment);
				//winners++;
			}
			delete disputes_[_id].votersIndex[disputes_[_id].voters[i].id];
		}
		// payment =  totalVoters_[_id]*voteCost_-winers*voteCost_;
		// dia_.transfer(msg.sender, payment);
		delete disputes_[_id];
		emit DisputeClosed(_id, drop);
	}

	/**
	* @dev get dispute status.
	* @param _id data source id.
	*/
	function isDisputeOpen(uint256 _id) public view returns (bool) {
		return (disputes_[_id].deadline>0);
	}

	/**
	* @dev check if address voted already.
	* @param _id data source identifier.
	*/
	function didCastVote(uint256 _id) public view returns (bool){
		return (disputes_[_id].votersIndex[msg.sender]>0);
	}
}