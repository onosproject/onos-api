# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/config/v2/configuration.proto, onos/config/v2/object.proto, onos/config/v2/transaction.proto, onos/config/v2/value.proto
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


class TransactionState(betterproto.Enum):
    """TransactionState is the transaction state of a transaction phase"""

    # TRANSACTION_PENDING indicates the transaction is pending
    TRANSACTION_PENDING = 0
    # TRANSACTION_COMPLETE indicates the transaction is complete
    TRANSACTION_COMPLETE = 2
    # TRANSACTION_FAILED indicates the transaction failed
    TRANSACTION_FAILED = 3
    # TRANSACTION_VALIDATING indicates the transaction is in the validating state
    TRANSACTION_VALIDATING = 4
    # TRANSACTION_APPLYING indicates the transaction is in the applying state
    TRANSACTION_APPLYING = 5


class TransactionEventType(betterproto.Enum):
    """TransactionEventType transaction event types for transaction store"""

    TRANSACTION_EVENT_UNKNOWN = 0
    TRANSACTION_CREATED = 1
    TRANSACTION_UPDATED = 2
    TRANSACTION_DELETED = 3
    TRANSACTION_REPLAYED = 4


class PathState(betterproto.Enum):
    """PathState is the state of a configuration path"""

    PATH_UPDATE_PENDING = 0
    PATH_UPDATE_COMPLETE = 1


class ConfigurationState(betterproto.Enum):
    """
    ConfigurationState is the configuration state of a configuration phase
    """

    # CONFIGURATION_PENDING indicates the configuration is PENDING
    CONFIGURATION_PENDING = 0
    # CONFIGURATION_SYNCHRONIZING indicates the configuration is synchronizing
    CONFIGURATION_SYNCHRONIZING = 1
    # CONFIGURATION_COMPLETE indicates the configuration is complete
    CONFIGURATION_COMPLETE = 2
    # CONFIGURATION_FAILED indicates the configuration is failed
    CONFIGURATION_FAILED = 3


class ConfigurationEventType(betterproto.Enum):
    """
    ConfigurationEventType configuration event types for configuration store
    """

    # CONFIGURATION_EVENT_UNKNOWN indicates unknown configuration store event
    CONFIGURATION_EVENT_UNKNOWN = 0
    # CONFIGURATION_CREATED indicates the configuration entry in the store is
    # created
    CONFIGURATION_CREATED = 1
    # CONFIGURATION_UPDATED indicates the configuration entry in the store is
    # updated
    CONFIGURATION_UPDATED = 2
    # CONFIGURATION_DELETED indicates the configuration entry in the store is
    # deleted
    CONFIGURATION_DELETED = 3
    # CONFIGURATION_REPLAYED
    CONFIGURATION_REPLAYED = 4


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
class PathValue(betterproto.Message):
    """PathValue is the state of a path/value in the configuration tree"""

    # 'path' is the path to change
    path: str = betterproto.string_field(1)
    # 'value' is the change value
    value: "TypedValue" = betterproto.message_field(2)
    # 'deleted' indicates whether this is a delete
    deleted: bool = betterproto.bool_field(3)
    # 'index'
    index: int = betterproto.uint64_field(4)


@dataclass(eq=False, repr=False)
class ObjectMeta(betterproto.Message):
    key: str = betterproto.string_field(1)
    version: int = betterproto.uint64_field(2)
    revision: int = betterproto.uint64_field(3)
    created: datetime = betterproto.message_field(4)
    updated: datetime = betterproto.message_field(5)
    deleted: datetime = betterproto.message_field(6)


@dataclass(eq=False, repr=False)
class TransactionChange(betterproto.Message):
    """TransactionChange  refers to a multi-target transactional change"""

    # 'changes' is a set of changes to apply to targets The list of changes
    # should contain only a single change per target/version pair.
    changes: Dict[str, "Change"] = betterproto.map_field(
        1, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class TransactionRollback(betterproto.Message):
    """TransactionRollback"""

    # 'index' is a monotonically increasing, globally unique index of the change
    index: int = betterproto.uint64_field(1)


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
    # 'status' is the current lifecycle status of the transaction
    status: "TransactionStatus" = betterproto.message_field(4)
    # 'username' is the name of the user that made the transaction
    username: str = betterproto.string_field(5)
    # atomic determines if a transaction is atomic or not
    atomic: bool = betterproto.bool_field(6)
    change: "TransactionChange" = betterproto.message_field(7, group="transaction")
    rollback: "TransactionRollback" = betterproto.message_field(8, group="transaction")


@dataclass(eq=False, repr=False)
class Change(betterproto.Message):
    """Change represents a configuration change to a single target"""

    # 'target_version' is an optional target version to which to apply this
    # change
    target_version: str = betterproto.string_field(2)
    # 'target_type' is an optional target type to which to apply this change
    target_type: str = betterproto.string_field(3)
    # 'values' is a set of change values to apply
    values: Dict[str, "ChangeValue"] = betterproto.map_field(
        4, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class ChangeValue(betterproto.Message):
    """
    ChangeValue represents a change requested for an individual path/value
    """

    # 'value' is the change value
    value: "TypedValue" = betterproto.message_field(2)
    # 'delete' indicates whether this is a delete
    delete: bool = betterproto.bool_field(3)


@dataclass(eq=False, repr=False)
class TransactionStatus(betterproto.Message):
    """TransactionStatus is the status of a Transaction"""

    # revision is the highest revision number that's been reconciled
    revision: int = betterproto.uint64_field(5)
    # 'state' is the state of the transaction This field should only be updated
    # from within onos-config.
    state: "TransactionState" = betterproto.enum_field(2)
    # 'sources' is a set of changes needed to revert back to the source of the
    # transaction This field should only be updated from within onos-config
    sources: Dict[str, "Source"] = betterproto.map_field(
        3, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class Source(betterproto.Message):
    """Source is a transaction source"""

    # 'target_version' is an optional target version to which to apply this
    # change
    target_version: str = betterproto.string_field(1)
    # 'target_type' is an optional target type to which to apply this change
    target_type: str = betterproto.string_field(2)
    # 'values' is the set of values for the source
    values: Dict[str, "PathValue"] = betterproto.map_field(
        3, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class TransactionEvent(betterproto.Message):
    """TransactionEvent transaction store event"""

    type: "TransactionEventType" = betterproto.enum_field(1)
    transaction: "Transaction" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class Configuration(betterproto.Message):
    """Configuration represents complete desired target configuration"""

    meta: "ObjectMeta" = betterproto.message_field(1)
    # 'id' is a unique configuration identifier
    id: str = betterproto.string_field(2)
    # 'target_id' is the target to which the desired target configuration applies
    target_id: str = betterproto.string_field(3)
    # 'target_version' is the version to which desired target configuration
    # applies
    target_version: str = betterproto.string_field(4)
    # 'target_type' is an optional target type to which to apply this desired
    # target configuration
    target_type: str = betterproto.string_field(5)
    # 'values' is a map of path/values to set
    values: Dict[str, "PathValue"] = betterproto.map_field(
        6, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )
    # 'ConfigurationStatus' is the current lifecycle status of the configuration
    status: "ConfigurationStatus" = betterproto.message_field(7)


@dataclass(eq=False, repr=False)
class ConfigurationStatus(betterproto.Message):
    """ConfigurationStatus is the status of a Configuration"""

    # revision is the highest revision number that's been reconciled
    revision: int = betterproto.uint64_field(1)
    # 'state' is the state of the transaction within a Phase
    state: "ConfigurationState" = betterproto.enum_field(2)
    # mastershipState mastership info
    mastership_state: "MastershipState" = betterproto.message_field(3)
    # paths a set of path statuses
    paths: Dict[str, "PathStatus"] = betterproto.map_field(
        4, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class PathStatus(betterproto.Message):
    """PathStatus is the status of a Configuration path"""

    state: "PathState" = betterproto.enum_field(1)
    update_index: int = betterproto.uint64_field(3)


@dataclass(eq=False, repr=False)
class MastershipState(betterproto.Message):
    """Mastership state"""

    term: int = betterproto.uint64_field(1)


@dataclass(eq=False, repr=False)
class ConfigurationEvent(betterproto.Message):
    """ConfigurationEvent configuration store event"""

    # ConfigurationEventType configuration event type
    type: "ConfigurationEventType" = betterproto.enum_field(1)
    configuration: "Configuration" = betterproto.message_field(2)
