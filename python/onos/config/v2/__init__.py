# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/config/v2/configuration.proto, onos/config/v2/failure.proto, onos/config/v2/object.proto, onos/config/v2/proposal.proto, onos/config/v2/transaction.proto, onos/config/v2/value.proto
# plugin: python-betterproto
from dataclasses import dataclass
from datetime import datetime
from typing import Dict, List

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase


class ValueType(betterproto.Enum):
    """ValueType is the type for a value"""

    EMPTY = 0
    STRING = 1
    INT = 2
    UINT = 3
    BOOL = 4
    DECIMAL = 5
    FLOAT = 6
    BYTES = 7
    LEAFLIST_STRING = 8
    LEAFLIST_INT = 9
    LEAFLIST_UINT = 10
    LEAFLIST_BOOL = 11
    LEAFLIST_DECIMAL = 12
    LEAFLIST_FLOAT = 13
    LEAFLIST_BYTES = 14


class FailureType(betterproto.Enum):
    UNKNOWN = 0
    CANCELED = 1
    NOT_FOUND = 2
    ALREADY_EXISTS = 3
    UNAUTHORIZED = 4
    FORBIDDEN = 5
    CONFLICT = 6
    INVALID = 7
    UNAVAILABLE = 8
    NOT_SUPPORTED = 9
    TIMEOUT = 10
    INTERNAL = 11


class TransactionStrategySynchronicity(betterproto.Enum):
    ASYNCHRONOUS = 0
    SYNCHRONOUS = 1


class TransactionStrategyIsolation(betterproto.Enum):
    DEFAULT = 0
    SERIALIZABLE = 1


class TransactionStatusState(betterproto.Enum):
    PENDING = 0
    VALIDATED = 1
    COMMITTED = 2
    APPLIED = 3
    FAILED = 4


class TransactionInitializePhaseState(betterproto.Enum):
    INITIALIZING = 0
    INITIALIZED = 1
    FAILED = 2


class TransactionValidatePhaseState(betterproto.Enum):
    VALIDATING = 0
    VALIDATED = 1
    FAILED = 2


class TransactionCommitPhaseState(betterproto.Enum):
    COMMITTING = 0
    COMMITTED = 1


class TransactionApplyPhaseState(betterproto.Enum):
    APPLYING = 0
    APPLIED = 1


class TransactionEventEventType(betterproto.Enum):
    UNKNOWN = 0
    CREATED = 1
    UPDATED = 2
    DELETED = 3
    REPLAYED = 4


class ConfigurationStatusState(betterproto.Enum):
    UNKNOWN = 0
    SYNCHRONIZING = 1
    SYNCHRONIZED = 2
    PERSISTED = 3


class ConfigurationEventEventType(betterproto.Enum):
    UNKNOWN = 0
    CREATED = 1
    UPDATED = 2
    DELETED = 3
    REPLAYED = 4


class ProposalInitializePhaseState(betterproto.Enum):
    INITIALIZING = 0
    INITIALIZED = 1


class ProposalValidatePhaseState(betterproto.Enum):
    VALIDATING = 0
    VALIDATED = 1
    FAILED = 2


class ProposalCommitPhaseState(betterproto.Enum):
    COMMITTING = 0
    COMMITTED = 1


class ProposalApplyPhaseState(betterproto.Enum):
    APPLYING = 0
    APPLIED = 1


class ProposalEventEventType(betterproto.Enum):
    UNKNOWN = 0
    CREATED = 1
    UPDATED = 2
    DELETED = 3
    REPLAYED = 4


@dataclass(eq=False, repr=False)
class TypedValue(betterproto.Message):
    """TypedValue is a value represented as a byte array"""

    # 'bytes' is the bytes array
    bytes_: bytes = betterproto.bytes_field(1)
    # 'type' is the value type
    type: "ValueType" = betterproto.enum_field(2)
    # 'type_opts' is a set of type options
    type_opts: List[int] = betterproto.int32_field(3)


@dataclass(eq=False, repr=False)
class PathValues(betterproto.Message):
    """PathValues is a set of path/value pairs"""

    # 'values' is a set of change values to apply
    values: Dict[str, "PathValue"] = betterproto.map_field(
        4, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class PathValue(betterproto.Message):
    """PathValue is the state of a path/value in the configuration tree"""

    # 'path' is the path to change
    path: str = betterproto.string_field(1)
    # 'value' is the change value
    value: "TypedValue" = betterproto.message_field(2)
    # 'deleted' indicates whether this is a delete
    deleted: bool = betterproto.bool_field(3)


@dataclass(eq=False, repr=False)
class ObjectMeta(betterproto.Message):
    key: str = betterproto.string_field(1)
    version: int = betterproto.uint64_field(2)
    revision: int = betterproto.uint64_field(3)
    created: datetime = betterproto.message_field(4)
    updated: datetime = betterproto.message_field(5)
    deleted: datetime = betterproto.message_field(6)


@dataclass(eq=False, repr=False)
class Failure(betterproto.Message):
    """Failure transaction failure type and description"""

    type: "FailureType" = betterproto.enum_field(1)
    description: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class TransactionInfo(betterproto.Message):
    """
    TransactionInfo is an extension providing information about the transaction
    to clients in responses.
    """

    id: str = betterproto.string_field(1)
    index: int = betterproto.uint64_field(2)


@dataclass(eq=False, repr=False)
class Transaction(betterproto.Message):
    """Transaction refers to a transaction change or transaction rollback"""

    meta: "ObjectMeta" = betterproto.message_field(1)
    # 'id' is the unique identifier of the transaction This field should be set
    # prior to persisting the object.
    id: str = betterproto.string_field(2)
    # 'index' is a monotonically increasing, globally unique index of the change
    # The index is provided by the store, is static and unique for each unique
    # change identifier, and should not be modified by client code.
    index: int = betterproto.uint64_field(3)
    # 'username' is the name of the user that made the transaction
    username: str = betterproto.string_field(4)
    # 'strategy' is the transaction strategy
    strategy: "TransactionStrategy" = betterproto.message_field(5)
    change: "ChangeTransaction" = betterproto.message_field(6, group="details")
    rollback: "RollbackTransaction" = betterproto.message_field(7, group="details")
    # 'status' is the current lifecycle status of the transaction
    status: "TransactionStatus" = betterproto.message_field(8)


@dataclass(eq=False, repr=False)
class TransactionStrategy(betterproto.Message):
    # 'synchronicity' indicates the transaction synchronicity level
    synchronicity: "TransactionStrategySynchronicity" = betterproto.enum_field(1)
    # 'isolation' indicates the transaction isolation level
    isolation: "TransactionStrategyIsolation" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class ChangeTransaction(betterproto.Message):
    # 'values' is a set of changes to apply to targets
    values: Dict[str, "PathValues"] = betterproto.map_field(
        1, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class RollbackTransaction(betterproto.Message):
    # 'rollback_index' is the index of the transaction to roll back
    rollback_index: int = betterproto.uint64_field(1)


@dataclass(eq=False, repr=False)
class TransactionStatus(betterproto.Message):
    # 'phases' is the transaction phases
    phases: "TransactionPhases" = betterproto.message_field(1)
    # 'proposals' is the set of proposals managed by the transaction
    proposals: List[str] = betterproto.string_field(2)
    # 'state' is the overall transaction state
    state: "TransactionStatusState" = betterproto.enum_field(3)
    # 'failure' is the transaction failure (if any)
    failure: "Failure" = betterproto.message_field(4)


@dataclass(eq=False, repr=False)
class TransactionPhases(betterproto.Message):
    # 'initialize' is the proposal initialization phase status
    initialize: "TransactionInitializePhase" = betterproto.message_field(1)
    # 'validate' is the proposal validation phase status
    validate: "TransactionValidatePhase" = betterproto.message_field(2)
    # 'commit' is the proposal commit phase status
    commit: "TransactionCommitPhase" = betterproto.message_field(3)
    # 'apply' is the proposal apply phase status
    apply: "TransactionApplyPhase" = betterproto.message_field(4)


@dataclass(eq=False, repr=False)
class TransactionPhaseStatus(betterproto.Message):
    start: datetime = betterproto.message_field(1)
    end: datetime = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class TransactionInitializePhase(betterproto.Message):
    status: "TransactionPhaseStatus" = betterproto.message_field(1)
    state: "TransactionInitializePhaseState" = betterproto.enum_field(2)
    failure: "Failure" = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class TransactionValidatePhase(betterproto.Message):
    status: "TransactionPhaseStatus" = betterproto.message_field(1)
    state: "TransactionValidatePhaseState" = betterproto.enum_field(2)
    failure: "Failure" = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class TransactionCommitPhase(betterproto.Message):
    status: "TransactionPhaseStatus" = betterproto.message_field(1)
    state: "TransactionCommitPhaseState" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class TransactionApplyPhase(betterproto.Message):
    status: "TransactionPhaseStatus" = betterproto.message_field(1)
    state: "TransactionApplyPhaseState" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class TransactionEvent(betterproto.Message):
    """TransactionEvent transaction store event"""

    type: "TransactionEventEventType" = betterproto.enum_field(1)
    transaction: "Transaction" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class Configuration(betterproto.Message):
    """Configuration represents complete desired target configuration"""

    meta: "ObjectMeta" = betterproto.message_field(1)
    # 'id' is a unique configuration identifier
    id: str = betterproto.string_field(2)
    # 'target_id' is the target to which the desired target configuration applies
    target_id: str = betterproto.string_field(3)
    # 'values' is a map of path/values to set
    values: Dict[str, "PathValue"] = betterproto.map_field(
        4, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )
    # 'index' is the index of the configuration values
    index: int = betterproto.uint64_field(5)
    # 'ConfigurationStatus' is the current lifecycle status of the configuration
    status: "ConfigurationStatus" = betterproto.message_field(6)


@dataclass(eq=False, repr=False)
class ConfigurationStatus(betterproto.Message):
    """ConfigurationStatus is the status of a Configuration"""

    # 'state' is the configuration state
    state: "ConfigurationStatusState" = betterproto.enum_field(1)
    # 'term' is the current mastership term for the configuration
    term: int = betterproto.uint64_field(2)
    # 'proposed' is the proposed configuration status
    proposed: "ProposedConfigurationStatus" = betterproto.message_field(3)
    # 'committed' is the committed configuration status
    committed: "CommittedConfigurationStatus" = betterproto.message_field(4)
    # 'applied' is the applied configuration status
    applied: "AppliedConfigurationStatus" = betterproto.message_field(5)


@dataclass(eq=False, repr=False)
class ProposedConfigurationStatus(betterproto.Message):
    index: int = betterproto.uint64_field(1)


@dataclass(eq=False, repr=False)
class CommittedConfigurationStatus(betterproto.Message):
    index: int = betterproto.uint64_field(1)


@dataclass(eq=False, repr=False)
class AppliedConfigurationStatus(betterproto.Message):
    index: int = betterproto.uint64_field(1)
    term: int = betterproto.uint64_field(2)
    values: Dict[str, "PathValue"] = betterproto.map_field(
        3, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class ConfigurationEvent(betterproto.Message):
    """ConfigurationEvent configuration store event"""

    # EventType configuration event type
    type: "ConfigurationEventEventType" = betterproto.enum_field(1)
    configuration: "Configuration" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class Proposal(betterproto.Message):
    meta: "ObjectMeta" = betterproto.message_field(1)
    # 'id' is the unique identifier of the proposal
    id: str = betterproto.string_field(2)
    # 'target_id' is the proposal's target identifier
    target_id: str = betterproto.string_field(3)
    # 'transaction_index' is the unique index of the transaction
    transaction_index: int = betterproto.uint64_field(4)
    change: "ChangeProposal" = betterproto.message_field(5, group="details")
    rollback: "RollbackProposal" = betterproto.message_field(6, group="details")
    # 'status' is the current lifecycle status of the proposal
    status: "ProposalStatus" = betterproto.message_field(7)


@dataclass(eq=False, repr=False)
class ChangeProposal(betterproto.Message):
    # 'changes' is the proposed change values
    values: Dict[str, "PathValue"] = betterproto.map_field(
        1, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class RollbackProposal(betterproto.Message):
    # 'rollback_index' is the index of the transaction to roll back
    rollback_index: int = betterproto.uint64_field(1)


@dataclass(eq=False, repr=False)
class ProposalStatus(betterproto.Message):
    """ProposalStatus is the status of a Proposal"""

    # 'phases' is the proposal phases
    phases: "ProposalPhases" = betterproto.message_field(1)
    # 'prev_index' is the index of the previous proposal associated with this
    # target
    prev_index: int = betterproto.uint64_field(2)
    # 'next_index' is the index of the next proposal associated with this target
    next_index: int = betterproto.uint64_field(3)
    # 'rollback_index' is a reference to the index to which to roll back
    rollback_index: int = betterproto.uint64_field(4)
    # 'rollback_values' is the set of values to use to roll back the proposal
    rollback_values: Dict[str, "PathValue"] = betterproto.map_field(
        5, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class ProposalPhases(betterproto.Message):
    # 'initialize' is the proposal initialization phase status
    initialize: "ProposalInitializePhase" = betterproto.message_field(1)
    # 'validate' is the proposal validation phase status
    validate: "ProposalValidatePhase" = betterproto.message_field(2)
    # 'commit' is the proposal commit phase status
    commit: "ProposalCommitPhase" = betterproto.message_field(3)
    # 'apply' is the proposal apply phase status
    apply: "ProposalApplyPhase" = betterproto.message_field(4)


@dataclass(eq=False, repr=False)
class ProposalPhaseStatus(betterproto.Message):
    start: datetime = betterproto.message_field(1)
    end: datetime = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class ProposalInitializePhase(betterproto.Message):
    status: "ProposalPhaseStatus" = betterproto.message_field(1)
    state: "ProposalInitializePhaseState" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class ProposalValidatePhase(betterproto.Message):
    status: "ProposalPhaseStatus" = betterproto.message_field(1)
    state: "ProposalValidatePhaseState" = betterproto.enum_field(2)
    failure: "Failure" = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class ProposalCommitPhase(betterproto.Message):
    status: "ProposalPhaseStatus" = betterproto.message_field(1)
    state: "ProposalCommitPhaseState" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class ProposalApplyPhase(betterproto.Message):
    status: "ProposalPhaseStatus" = betterproto.message_field(1)
    state: "ProposalApplyPhaseState" = betterproto.enum_field(2)
    term: int = betterproto.uint64_field(3)


@dataclass(eq=False, repr=False)
class ProposalEvent(betterproto.Message):
    """ProposalEvent proposal store event"""

    type: "ProposalEventEventType" = betterproto.enum_field(1)
    proposal: "Proposal" = betterproto.message_field(2)
