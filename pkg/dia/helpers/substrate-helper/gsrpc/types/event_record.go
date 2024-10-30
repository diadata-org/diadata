// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	"github.com/ethereum/go-ethereum/log"
)

// EventRecordsRaw is a raw record for a set of events, represented as the raw bytes. It exists since
// decoding of events can only be done with metadata, so events can't follow the static way of decoding
// other types do. It exposes functions to decode events using metadata and targets.
// Be careful using this in your own structs – it only works as the last value in a struct since it will consume the
// remainder of the encoded data. The reason for this is that it does not contain any length encoding, so it would
// not know where to stop.
//
// Deprecated: EventRecordsRaw relies on static event definition that is no longer maintained,
// please check retriever.EventRetriever.
type EventRecordsRaw []byte

// Encode implements encoding for Data, which just unwraps the bytes of Data
func (e EventRecordsRaw) Encode(encoder scale.Encoder) error {
	return encoder.Write(e)
}

// Decode implements decoding for Data, which just reads all the remaining bytes into Data
func (e *EventRecordsRaw) Decode(decoder scale.Decoder) error {
	for i := 0; true; i++ {
		b, err := decoder.ReadOneByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		*e = append((*e)[:i], b)
	}
	return nil
}

// EventRecords is a default set of possible event records that can be used as a target for
// `func (e EventRecordsRaw) Decode(...`
// Sources:
// https://github.com/polkadot-js/api/blob/master/packages/api-augment/src/substrate/events.ts
// https://github.com/polkadot-js/api/blob/master/packages/api-augment/src/polkadot/events.ts
//
//nolint:stylecheck,lll,revive
type EventRecords struct {
	Auctions_AuctionStarted     []EventAuctionsAuctionStarted
	Auctions_AuctionClosed      []EventAuctionsAuctionClosed
	Auctions_Reserved           []EventAuctionsReserved
	Auctions_Unreserved         []EventAuctionsUnreserved
	Auctions_ReserveConfiscated []EventAuctionsReserveConfiscated
	Auctions_BidAccepted        []EventAuctionsBidAccepted
	Auctions_WinningOffset      []EventAuctionsWinningOffset

	Assets_Created             []EventAssetCreated
	Assets_Issued              []EventAssetIssued
	Assets_Transferred         []EventAssetTransferred
	Assets_Burned              []EventAssetBurned
	Assets_TeamChanged         []EventAssetTeamChanged
	Assets_OwnerChanged        []EventAssetOwnerChanged
	Assets_Frozen              []EventAssetFrozen
	Assets_Thawed              []EventAssetThawed
	Assets_AssetFrozen         []EventAssetAssetFrozen
	Assets_AssetThawed         []EventAssetAssetThawed
	Assets_Destroyed           []EventAssetDestroyed
	Assets_ForceCreated        []EventAssetForceCreated
	Assets_MetadataSet         []EventAssetMetadataSet
	Assets_MetadataCleared     []EventAssetMetadataCleared
	Assets_ApprovedTransfer    []EventAssetApprovedTransfer
	Assets_ApprovalCancelled   []EventAssetApprovalCancelled
	Assets_TransferredApproved []EventAssetTransferredApproved
	Assets_AssetStatusChanged  []EventAssetAssetStatusChanged

	BagsList_Rebagged []EventBagsListRebagged

	Balances_BalanceSet         []EventBalancesBalanceSet
	Balances_Deposit            []EventBalancesDeposit
	Balances_DustLost           []EventBalancesDustLost
	Balances_Endowed            []EventBalancesEndowed
	Balances_Reserved           []EventBalancesReserved
	Balances_ReserveRepatriated []EventBalancesReserveRepatriated
	Balances_Slashed            []EventBalancesSlashed
	Balances_Transfer           []EventBalancesTransfer
	Balances_Unreserved         []EventBalancesUnreserved
	Balances_Withdraw           []EventBalancesWithdraw

	Bounties_BountyProposed     []EventBountiesBountyProposed
	Bounties_BountyRejected     []EventBountiesBountyRejected
	Bounties_BountyBecameActive []EventBountiesBountyBecameActive
	Bounties_BountyAwarded      []EventBountiesBountyAwarded
	Bounties_BountyClaimed      []EventBountiesBountyClaimed
	Bounties_BountyCanceled     []EventBountiesBountyCanceled
	Bounties_BountyExtended     []EventBountiesBountyExtended

	ChildBounties_Added    []EventChildBountiesAdded
	ChildBounties_Awarded  []EventChildBountiesAwarded
	ChildBounties_Claimed  []EventChildBountiesClaimed
	ChildBounties_Canceled []EventChildBountiesCanceled

	Claims_Claimed []EventClaimsClaimed

	CollatorSelection_NewInvulnerables     []EventCollatorSelectionNewInvulnerables
	CollatorSelection_NewDesiredCandidates []EventCollatorSelectionNewDesiredCandidates
	CollatorSelection_NewCandidacyBond     []EventCollatorSelectionNewCandidacyBond
	CollatorSelection_CandidateAdded       []EventCollatorSelectionCandidateAdded
	CollatorSelection_CandidateRemoved     []EventCollatorSelectionCandidateRemoved

	Contracts_CodeRemoved         []EventContractsCodeRemoved
	Contracts_CodeStored          []EventContractsCodeStored
	Contracts_ContractCodeUpdated []EventContractsContractCodeUpdated
	Contracts_ContractEmitted     []EventContractsContractEmitted
	Contracts_Instantiated        []EventContractsInstantiated
	Contracts_Terminated          []EventContractsTerminated

	ConvictionVoting_Delegated   []EventConvictionVotingDelegated
	ConvictionVoting_Undelegated []EventConvictionVotingUndelegated

	Council_Approved       []EventCouncilApproved
	Council_Closed         []EventCouncilClosed
	Council_Disapproved    []EventCouncilDisapproved
	Council_Executed       []EventCouncilExecuted
	Council_MemberExecuted []EventCouncilMemberExecuted
	Council_Proposed       []EventCouncilProposed
	Council_Voted          []EventCouncilVoted

	Crowdloan_Created           []EventCrowdloanCreated
	Crowdloan_Contributed       []EventCrowdloanContributed
	Crowdloan_Withdrew          []EventCrowdloanWithdrew
	Crowdloan_PartiallyRefunded []EventCrowdloanPartiallyRefunded
	Crowdloan_AllRefunded       []EventCrowdloanAllRefunded
	Crowdloan_Dissolved         []EventCrowdloanDissolved
	Crowdloan_HandleBidResult   []EventCrowdloanHandleBidResult
	Crowdloan_Edited            []EventCrowdloanEdited
	Crowdloan_MemoUpdated       []EventCrowdloanMemoUpdated
	Crowdloan_AddedToNewRaise   []EventCrowdloanAddedToNewRaise

	Democracy_Blacklisted     []EventDemocracyBlacklisted
	Democracy_Cancelled       []EventDemocracyCancelled
	Democracy_Delegated       []EventDemocracyDelegated
	Democracy_Executed        []EventDemocracyExecuted
	Democracy_ExternalTabled  []EventDemocracyExternalTabled
	Democracy_NotPassed       []EventDemocracyNotPassed
	Democracy_Passed          []EventDemocracyPassed
	Democracy_PreimageInvalid []EventDemocracyPreimageInvalid
	Democracy_PreimageMissing []EventDemocracyPreimageMissing
	Democracy_PreimageNoted   []EventDemocracyPreimageNoted
	Democracy_PreimageReaped  []EventDemocracyPreimageReaped
	Democracy_PreimageUsed    []EventDemocracyPreimageUsed
	Democracy_Proposed        []EventDemocracyProposed
	Democracy_Seconded        []EventDemocracySeconded
	Democracy_Started         []EventDemocracyStarted
	Democracy_Tabled          []EventDemocracyTabled
	Democracy_Undelegated     []EventDemocracyUndelegated
	Democracy_Vetoed          []EventDemocracyVetoed
	Democracy_Voted           []EventDemocracyVoted

	ElectionProviderMultiPhase_SolutionStored       []EventElectionProviderMultiPhaseSolutionStored
	ElectionProviderMultiPhase_ElectionFinalized    []EventElectionProviderMultiPhaseElectionFinalized
	ElectionProviderMultiPhase_Rewarded             []EventElectionProviderMultiPhaseRewarded
	ElectionProviderMultiPhase_Slashed              []EventElectionProviderMultiPhaseSlashed
	ElectionProviderMultiPhase_SignedPhaseStarted   []EventElectionProviderMultiPhaseSignedPhaseStarted
	ElectionProviderMultiPhase_UnsignedPhaseStarted []EventElectionProviderMultiPhaseUnsignedPhaseStarted

	Elections_CandidateSlashed  []EventElectionsCandidateSlashed
	Elections_ElectionError     []EventElectionsElectionError
	Elections_EmptyTerm         []EventElectionsEmptyTerm
	Elections_MemberKicked      []EventElectionsMemberKicked
	Elections_NewTerm           []EventElectionsNewTerm
	Elections_Renounced         []EventElectionsRenounced
	Elections_SeatHolderSlashed []EventElectionsSeatHolderSlashed

	Gilt_BidPlaced    []EventGiltBidPlaced
	Gilt_BidRetracted []EventGiltBidRetracted
	Gilt_GiltIssued   []EventGiltGiltIssued
	Gilt_GiltThawed   []EventGiltGiltThawed

	Grandpa_NewAuthorities []EventGrandpaNewAuthorities
	Grandpa_Paused         []EventGrandpaPaused
	Grandpa_Resumed        []EventGrandpaResumed

	Hrmp_OpenChannelRequested []EventHRMPOpenChannelRequested
	Hrmp_OpenChannelCanceled  []EventHRMPOpenChannelCanceled
	Hrmp_OpenChannelAccepted  []EventHRMPOpenChannelAccepted
	Hrmp_ChannelClosed        []EventHRMPChannelClosed

	Identity_IdentityCleared      []EventIdentityCleared
	Identity_IdentityKilled       []EventIdentityKilled
	Identity_IdentitySet          []EventIdentitySet
	Identity_JudgementGiven       []EventIdentityJudgementGiven
	Identity_JudgementRequested   []EventIdentityJudgementRequested
	Identity_JudgementUnrequested []EventIdentityJudgementUnrequested
	Identity_RegistrarAdded       []EventIdentityRegistrarAdded
	Identity_SubIdentityAdded     []EventIdentitySubIdentityAdded
	Identity_SubIdentityRemoved   []EventIdentitySubIdentityRemoved
	Identity_SubIdentityRevoked   []EventIdentitySubIdentityRevoked

	ImOnline_AllGood           []EventImOnlineAllGood
	ImOnline_HeartbeatReceived []EventImOnlineHeartbeatReceived
	ImOnline_SomeOffline       []EventImOnlineSomeOffline

	Indices_IndexAssigned []EventIndicesIndexAssigned
	Indices_IndexFreed    []EventIndicesIndexFreed
	Indices_IndexFrozen   []EventIndicesIndexFrozen

	Lottery_LotteryStarted []EventLotteryLotteryStarted
	Lottery_CallsUpdated   []EventLotteryCallsUpdated
	Lottery_Winner         []EventLotteryWinner
	Lottery_TicketBought   []EventLotteryTicketBought

	Multisig_MultisigApproval  []EventMultisigApproval
	Multisig_MultisigCancelled []EventMultisigCancelled
	Multisig_MultisigExecuted  []EventMultisigExecuted
	Multisig_NewMultisig       []EventMultisigNewMultisig

	NftSales_ForSale []EventNftSalesForSale
	NftSales_Removed []EventNftSalesRemoved
	NftSales_Sold    []EventNftSalesSold

	Offences_Offence []EventOffencesOffence

	OrmlAssetRegistry_RegisteredAsset []EventOrmlAssetRegistryRegisteredAsset
	OrmlAssetRegistry_UpdatedAsset    []EventOrmlAssetRegistryUpdatedAsset

	OrmlTokens_Endowed            []EventOrmlTokensEndowed
	OrmlTokens_DustLost           []EventOrmlTokensDustLost
	OrmlTokens_Transfer           []EventOrmlTokensTransfer
	OrmlTokens_Reserved           []EventOrmlTokensReserved
	OrmlTokens_Unreserved         []EventOrmlTokensUnreserved
	OrmlTokens_ReserveRepatriated []EventOrmlTokensReserveRepatriated
	OrmlTokens_BalanceSet         []EventOrmlTokensBalanceSet
	OrmlTokens_TotalIssuanceSet   []EventOrmlTokensTotalIssuanceSet
	OrmlTokens_Withdrawn          []EventOrmlTokensWithdrawn
	OrmlTokens_Slashed            []EventOrmlTokensSlashed
	OrmlTokens_Deposited          []EventOrmlTokensDeposited
	OrmlTokens_LockSet            []EventOrmlTokensLockSet
	OrmlTokens_LockRemoved        []EventOrmlTokensLockRemoved
	OrmlTokens_Locked             []EventOrmlTokensLocked
	OrmlTokens_Unlocked           []EventOrmlTokensUnlocked

	Paras_CurrentCodeUpdated   []EventParasCurrentCodeUpdated
	Paras_CurrentHeadUpdated   []EventParasCurrentHeadUpdated
	Paras_CodeUpgradeScheduled []EventParasCodeUpgradeScheduled
	Paras_NewHeadNoted         []EventParasNewHeadNoted
	Paras_ActionQueued         []EventParasActionQueued
	Paras_PvfCheckStarted      []EventParasPvfCheckStarted
	Paras_PvfCheckAccepted     []EventParasPvfCheckAccepted
	Paras_PvfCheckRejected     []EventParasPvfCheckRejected

	ParasDisputes_DisputeInitiated []EventParasDisputesDisputeInitiated
	ParasDisputes_DisputeConcluded []EventParasDisputesDisputeConcluded
	ParasDisputes_DisputeTimedOut  []EventParasDisputesDisputeTimedOut
	ParasDisputes_Revert           []EventParasDisputesRevert

	ParaInclusion_CandidateBacked   []EventParaInclusionCandidateBacked
	ParaInclusion_CandidateIncluded []EventParaInclusionCandidateIncluded
	ParaInclusion_CandidateTimedOut []EventParaInclusionCandidateTimedOut

	ParachainSystem_ValidationFunctionStored    []EventParachainSystemValidationFunctionStored
	ParachainSystem_ValidationFunctionApplied   []EventParachainSystemValidationFunctionApplied
	ParachainSystem_ValidationFunctionDiscarded []EventParachainSystemValidationFunctionDiscarded
	ParachainSystem_UpgradeAuthorized           []EventParachainSystemUpgradeAuthorized
	ParachainSystem_DownwardMessagesReceived    []EventParachainSystemDownwardMessagesReceived
	ParachainSystem_DownwardMessagesProcessed   []EventParachainSystemDownwardMessagesProcessed

	Preimage_Cleared   []EventPreimageCleared
	Preimage_Noted     []EventPreimageNoted
	Preimage_Requested []EventPreimageRequested

	Proxy_Announced     []EventProxyAnnounced
	Proxy_PureCreated   []EventProxyPureCreated
	Proxy_ProxyAdded    []EventProxyProxyAdded
	Proxy_ProxyExecuted []EventProxyProxyExecuted
	Proxy_ProxyRemoved  []EventProxyProxyRemoved

	Recovery_AccountRecovered  []EventRecoveryAccountRecovered
	Recovery_RecoveryClosed    []EventRecoveryClosed
	Recovery_RecoveryCreated   []EventRecoveryCreated
	Recovery_RecoveryInitiated []EventRecoveryInitiated
	Recovery_RecoveryRemoved   []EventRecoveryRemoved
	Recovery_RecoveryVouched   []EventRecoveryVouched

	Registrar_Registered   []EventRegistrarRegistered
	Registrar_Deregistered []EventRegistrarDeregistered
	Registrar_Reserved     []EventRegistrarReserved

	Referenda_Submitted               []EventReferendaSubmitted
	Referenda_DecisionDepositPlaced   []EventReferendaDecisionDepositPlaced
	Referenda_DecisionDepositRefunded []EventReferendaDecisionDepositRefunded
	Referenda_DepositSlashed          []EventReferendaDecisionSlashed
	Referenda_DecisionStarted         []EventReferendaDecisionStarted
	Referenda_ConfirmStarted          []EventReferendaConfirmStarted
	Referenda_ConfirmAborted          []EventReferendaConfirmAborted
	Referenda_Confirmed               []EventReferendaConfirmed
	Referenda_Approved                []EventReferendaApproved
	Referenda_Rejected                []EventReferendaRejected
	Referenda_TimedOut                []EventReferendaTimedOut
	Referenda_Cancelled               []EventReferendaCancelled
	Referenda_Killed                  []EventReferendaKilled

	Scheduler_CallLookupFailed []EventSchedulerCallLookupFailed
	Scheduler_Canceled         []EventSchedulerCanceled
	Scheduler_Dispatched       []EventSchedulerDispatched
	Scheduler_Scheduled        []EventSchedulerScheduled

	Session_NewSession []EventSessionNewSession

	Slots_NewLeasePeriod []EventSlotsNewLeasePeriod
	Slots_Leased         []EventSlotsLeased

	Society_AutoUnbid                []EventSocietyAutoUnbid
	Society_Bid                      []EventSocietyBid
	Society_CandidateSuspended       []EventSocietyCandidateSuspended
	Society_Challenged               []EventSocietyChallenged
	Society_DefenderVote             []EventSocietyDefenderVote
	Society_Deposit                  []EventSocietyDeposit
	Society_Founded                  []EventSocietyFounded
	Society_Inducted                 []EventSocietyInducted
	Society_MemberSuspended          []EventSocietyMemberSuspended
	Society_NewMaxMembers            []EventSocietyNewMaxMembers
	Society_SuspendedMemberJudgement []EventSocietySuspendedMemberJudgement
	Society_Unbid                    []EventSocietyUnbid
	Society_Unfounded                []EventSocietyUnfounded
	Society_Unvouch                  []EventSocietyUnvouch
	Society_Vote                     []EventSocietyVote
	Society_Vouch                    []EventSocietyVouch

	Staking_Bonded                     []EventStakingBonded
	Staking_Chilled                    []EventStakingChilled
	Staking_EraPaid                    []EventStakingEraPaid
	Staking_Kicked                     []EventStakingKicked
	Staking_OldSlashingReportDiscarded []EventStakingOldSlashingReportDiscarded
	Staking_PayoutStarted              []EventStakingPayoutStarted
	Staking_Rewarded                   []EventStakingRewarded
	Staking_Slashed                    []EventStakingSlashed
	Staking_StakersElected             []EventStakingStakersElected
	Staking_StakingElectionFailed      []EventStakingStakingElectionFailed
	Staking_Unbonded                   []EventStakingUnbonded
	Staking_Withdrawn                  []EventStakingWithdrawn

	StateTrieMigration_Migrated              []EventStateTrieMigrationMigrated
	StateTrieMigration_Slashed               []EventStateTrieMigrationSlashed
	StateTrieMigration_AutoMigrationFinished []EventStateTrieMigrationAutoMigrationFinished
	StateTrieMigration_Halted                []EventStateTrieMigrationHalted

	Sudo_KeyChanged []EventSudoKeyChanged
	Sudo_Sudid      []EventSudoSudid
	Sudo_SudoAsDone []EventSudoAsDone

	System_CodeUpdated      []EventSystemCodeUpdated
	System_ExtrinsicFailed  []EventSystemExtrinsicFailed
	System_ExtrinsicSuccess []EventSystemExtrinsicSuccess
	System_KilledAccount    []EventSystemKilledAccount
	System_NewAccount       []EventSystemNewAccount
	System_Remarked         []EventSystemRemarked

	TechnicalCommittee_Approved       []EventTechnicalCommitteeApproved
	TechnicalCommittee_Closed         []EventTechnicalCommitteeClosed
	TechnicalCommittee_Disapproved    []EventTechnicalCommitteeDisapproved
	TechnicalCommittee_Executed       []EventTechnicalCommitteeExecuted
	TechnicalCommittee_MemberExecuted []EventTechnicalCommitteeMemberExecuted
	TechnicalCommittee_Proposed       []EventTechnicalCommitteeProposed
	TechnicalCommittee_Voted          []EventTechnicalCommitteeVoted

	TechnicalMembership_Dummy          []EventTechnicalMembershipDummy
	TechnicalMembership_KeyChanged     []EventTechnicalMembershipKeyChanged
	TechnicalMembership_MemberAdded    []EventTechnicalMembershipMemberAdded
	TechnicalMembership_MemberRemoved  []EventTechnicalMembershipMemberRemoved
	TechnicalMembership_MembersReset   []EventTechnicalMembershipMembersReset
	TechnicalMembership_MembersSwapped []EventTechnicalMembershipMembersSwapped

	Tips_NewTip       []EventTipsNewTip
	Tips_TipClosed    []EventTipsTipClosed
	Tips_TipClosing   []EventTipsTipClosing
	Tips_TipRetracted []EventTipsTipRetracted
	Tips_TipSlashed   []EventTipsTipSlashed

	TransactionStorage_Stored       []EventTransactionStorageStored
	TransactionStorage_Renewed      []EventTransactionStorageRenewed
	TransactionStorage_ProofChecked []EventTransactionStorageProofChecked

	TransactionPayment_TransactionFeePaid []EventTransactionPaymentTransactionFeePaid

	Treasury_Proposed        []EventTreasuryProposed
	Treasury_Spending        []EventTreasurySpending
	Treasury_Awarded         []EventTreasuryAwarded
	Treasury_Rejected        []EventTreasuryRejected
	Treasury_Burnt           []EventTreasuryBurnt
	Treasury_Rollover        []EventTreasuryRollover
	Treasury_Deposit         []EventTreasuryDeposit
	Treasury_SpendApproved   []EventTreasurySpendApproved
	Treasury_UpdatedInactive []EventTreasuryUpdatedInactive

	Uniques_ApprovalCancelled    []EventUniquesApprovalCancelled
	Uniques_ApprovedTransfer     []EventUniquesApprovedTransfer
	Uniques_AssetStatusChanged   []EventUniquesAssetStatusChanged
	Uniques_AttributeCleared     []EventUniquesAttributeCleared
	Uniques_AttributeSet         []EventUniquesAttributeSet
	Uniques_Burned               []EventUniquesBurned
	Uniques_ClassFrozen          []EventUniquesClassFrozen
	Uniques_ClassMetadataCleared []EventUniquesClassMetadataCleared
	Uniques_ClassMetadataSet     []EventUniquesClassMetadataSet
	Uniques_ClassThawed          []EventUniquesClassThawed
	Uniques_Created              []EventUniquesCreated
	Uniques_Destroyed            []EventUniquesDestroyed
	Uniques_ForceCreated         []EventUniquesForceCreated
	Uniques_Frozen               []EventUniquesFrozen
	Uniques_Issued               []EventUniquesIssued
	Uniques_MetadataCleared      []EventUniquesMetadataCleared
	Uniques_MetadataSet          []EventUniquesMetadataSet
	Uniques_OwnerChanged         []EventUniquesOwnerChanged
	Uniques_Redeposited          []EventUniquesRedeposited
	Uniques_TeamChanged          []EventUniquesTeamChanged
	Uniques_Thawed               []EventUniquesThawed
	Uniques_Transferred          []EventUniquesTransferred

	Ump_InvalidFormat          []EventUMPInvalidFormat
	Ump_UnsupportedVersion     []EventUMPUnsupportedVersion
	Ump_ExecutedUpward         []EventUMPExecutedUpward
	Ump_WeightExhausted        []EventUMPWeightExhausted
	Ump_UpwardMessagesReceived []EventUMPUpwardMessagesReceived
	Ump_OverweightEnqueued     []EventUMPOverweightEnqueued
	Ump_OverweightServiced     []EventUMPOverweightServiced

	Utility_BatchCompleted   []EventUtilityBatchCompleted
	Utility_BatchInterrupted []EventUtilityBatchInterrupted
	Utility_DispatchedAs     []EventUtilityBatchInterrupted
	Utility_ItemCompleted    []EventUtilityItemCompleted

	Vesting_VestingCompleted []EventVestingVestingCompleted
	Vesting_VestingUpdated   []EventVestingVestingUpdated

	VoterList_Rebagged     []EventVoterListRebagged
	VoterList_ScoreUpdated []EventVoterListScoreUpdated

	Whitelist_CallWhitelisted           []EventWhitelistCallWhitelisted
	Whitelist_WhitelistedCallRemoved    []EventWhitelistWhitelistedCallRemoved
	Whitelist_WhitelistedCallDispatched []EventWhitelistWhitelistedCallRemoved

	XcmPallet_Attempted                 []EventXcmPalletAttempted
	XcmPallet_Sent                      []EventXcmPalletSent
	XcmPallet_UnexpectedResponse        []EventXcmPalletUnexpectedResponse
	XcmPallet_ResponseReady             []EventXcmPalletResponseReady
	XcmPallet_Notified                  []EventXcmPalletNotified
	XcmPallet_NotifyOverweight          []EventXcmPalletNotifyOverweight
	XcmPallet_NotifyDispatchError       []EventXcmPalletNotifyDispatchError
	XcmPallet_NotifyDecodeFailed        []EventXcmPalletNotifyDecodeFailed
	XcmPallet_InvalidResponder          []EventXcmPalletInvalidResponder
	XcmPallet_InvalidResponderVersion   []EventXcmPalletInvalidResponderVersion
	XcmPallet_ResponseTaken             []EventXcmPalletResponseTaken
	XcmPallet_AssetsTrapped             []EventXcmPalletAssetsTrapped
	XcmPallet_VersionChangeNotified     []EventXcmPalletVersionChangeNotified
	XcmPallet_SupportedVersionChanged   []EventXcmPalletSupportedVersionChanged
	XcmPallet_NotifyTargetSendFail      []EventXcmPalletNotifyTargetSendFail
	XcmPallet_NotifyTargetMigrationFail []EventXcmPalletNotifyTargetMigrationFail
}

// DecodeEventRecords decodes the events records from an EventRecordRaw into a target t using the given Metadata m
// If this method returns an error like `unable to decode Phase for event #x: EOF`, it is likely that you have defined
// a custom event record with a wrong type. For example your custom event record has a field with a length prefixed
// type, such as types.Bytes, where your event in reallity contains a fixed width type, such as a types.U32.
func (e EventRecordsRaw) DecodeEventRecords(m *Metadata, t interface{}) error { //nolint:funlen
	log.Debug(fmt.Sprintf("will decode event records from raw hex: %#x", e))

	// ensure t is a pointer
	ttyp := reflect.TypeOf(t)
	if ttyp.Kind() != reflect.Ptr {
		return errors.New("target must be a pointer, but is " + fmt.Sprint(ttyp))
	}
	// ensure t is not a nil pointer
	tval := reflect.ValueOf(t)
	if tval.IsNil() {
		return errors.New("target is a nil pointer")
	}
	val := tval.Elem()
	typ := val.Type()
	// ensure val can be set
	if !val.CanSet() {
		return fmt.Errorf("unsettable value %v", typ)
	}
	// ensure val points to a struct
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("target must point to a struct, but is " + fmt.Sprint(typ))
	}

	decoder := scale.NewDecoder(bytes.NewReader(e))

	// determine number of events
	n, err := decoder.DecodeUintCompact()
	if err != nil {
		return err
	}

	log.Debug(fmt.Sprintf("found %v events", n))

	// iterate over events
	for i := uint64(0); i < n.Uint64(); i++ {
		log.Debug(fmt.Sprintf("decoding event #%v", i))

		// decode Phase
		phase := Phase{}
		err := decoder.Decode(&phase)
		if err != nil {
			return fmt.Errorf("unable to decode Phase for event #%v: %v", i, err)
		}

		// decode EventID
		id := EventID{}
		err = decoder.Decode(&id)
		if err != nil {
			return fmt.Errorf("unable to decode EventID for event #%v: %v", i, err)
		}

		log.Debug(fmt.Sprintf("event #%v has EventID %v", i, id))

		// ask metadata for method & event name for event
		moduleName, eventName, err := m.FindEventNamesForEventID(id)
		// moduleName, eventName, err := "System", "ExtrinsicSuccess", nil
		if err != nil {
			return fmt.Errorf("unable to find event with EventID %v in metadata for event #%v: %s", id, i, err)
		}

		log.Debug(fmt.Sprintf("event #%v is in module %v with event name %v", i, moduleName, eventName))

		// check whether name for eventID exists in t
		field := val.FieldByName(fmt.Sprintf("%v_%v", moduleName, eventName))
		if !field.IsValid() {
			return fmt.Errorf("unable to find field %v_%v for event #%v with EventID %v", moduleName, eventName, i, id)
		}

		// create a pointer to with the correct type that will hold the decoded event
		holder := reflect.New(field.Type().Elem())

		// ensure first field is for Phase, last field is for Topics
		numFields := holder.Elem().NumField()
		if numFields < 2 {
			return fmt.Errorf("expected event #%v with EventID %v, field %v_%v to have at least 2 fields "+
				"(for Phase and Topics), but has %v fields", i, id, moduleName, eventName, numFields)
		}
		phaseField := holder.Elem().FieldByIndex([]int{0})
		if phaseField.Type() != reflect.TypeOf(phase) {
			return fmt.Errorf("expected the first field of event #%v with EventID %v, field %v_%v to be of type "+
				"types.Phase, but got %v", i, id, moduleName, eventName, phaseField.Type())
		}
		topicsField := holder.Elem().FieldByIndex([]int{numFields - 1})
		if topicsField.Type() != reflect.TypeOf([]Hash{}) {
			return fmt.Errorf("expected the last field of event #%v with EventID %v, field %v_%v to be of type "+
				"[]types.Hash for Topics, but got %v", i, id, moduleName, eventName, topicsField.Type())
		}

		// set the phase we decoded earlier
		phaseField.Set(reflect.ValueOf(phase))

		// set the remaining fields
		for j := 1; j < numFields; j++ {
			err = decoder.Decode(holder.Elem().FieldByIndex([]int{j}).Addr().Interface())
			if err != nil {
				return fmt.Errorf("unable to decode field %v event #%v with EventID %v, field %v_%v: %v", j, i, id, moduleName,
					eventName, err)
			}
		}

		// add the decoded event to the slice
		field.Set(reflect.Append(field, holder.Elem()))

		log.Debug(fmt.Sprintf("decoded event #%v", i))
	}
	return nil
}

// Phase is an enum describing the current phase of the event (applying the extrinsic or finalized)
type Phase struct {
	IsApplyExtrinsic bool
	AsApplyExtrinsic uint32
	IsFinalization   bool
	IsInitialization bool
}

func (p *Phase) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsApplyExtrinsic = true
		err = decoder.Decode(&p.AsApplyExtrinsic)
	case 1:
		p.IsFinalization = true
	case 2:
		p.IsInitialization = true
	}

	if err != nil {
		return err
	}

	return nil
}

func (p Phase) Encode(encoder scale.Encoder) error {
	var err1, err2 error

	switch {
	case p.IsApplyExtrinsic:
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(p.AsApplyExtrinsic)
	case p.IsFinalization:
		err1 = encoder.PushByte(1)
	case p.IsInitialization:
		err1 = encoder.PushByte(2)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

type EventID [2]byte
