# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/config/admin/admin.proto
# plugin: python-betterproto
from dataclasses import dataclass
from datetime import timedelta
from typing import AsyncIterable, AsyncIterator, Dict, Iterable, List, Union

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase
import grpclib


class Type(betterproto.Enum):
    """Streaming event type"""

    # NONE indicates this response does not represent a state change
    NONE = 0
    # ADDED is an event which occurs when an item is added
    ADDED = 1
    # UPDATED is an event which occurs when an item is updated
    UPDATED = 2
    # REMOVED is an event which occurs when an item is removed
    REMOVED = 3


@dataclass(eq=False, repr=False)
class ReadOnlySubPath(betterproto.Message):
    """
    ReadOnlySubPath is an extension to the ReadOnlyPath to define the datatype
    of the subpath
    """

    # sub_path is the relative path of a child object e.g. /list2b/index
    sub_path: str = betterproto.string_field(1)
    # value_type is the datatype of the read only path
    value_type: "_v2__.ValueType" = betterproto.enum_field(2)
    type_opts: List[int] = betterproto.uint64_field(3)
    description: str = betterproto.string_field(4)
    units: str = betterproto.string_field(5)
    is_a_key: bool = betterproto.bool_field(6)
    attr_name: str = betterproto.string_field(7)


@dataclass(eq=False, repr=False)
class ReadOnlyPath(betterproto.Message):
    """
    ReadOnlyPath extracted from the model plugin as the definition of a tree of
    read only items. In YANG models items are defined as ReadOnly with the
    `config false` keyword. This can be applied to single items (leafs) or
    collections (containers or lists). When this `config false` is applied to
    an object every item beneath it will also become readonly - here these are
    shown as subpaths. The complete read only path then will be a concatenation
    of both e.g. /cont1a/cont1b-state/list2b/index and the type is defined in
    the SubPath as UInt8.
    """

    # path of the topmost `config false` object e.g. /cont1a/cont1b-state
    path: str = betterproto.string_field(1)
    # ReadOnlySubPath is a set of children of the path including an entry for the
    # type of the topmost object with subpath `/` An example is /list2b/index
    sub_path: List["ReadOnlySubPath"] = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class ReadWritePath(betterproto.Message):
    """
    ReadWritePath is extracted from the model plugin as the definition of a
    writeable attributes. In YANG models items are writable by default unless
    they are specified as `config false` or have an item with `config false` as
    a parent. Each configurable item has metadata with meanings taken from the
    YANG specification RFC 6020.
    """

    # path is the full path to the attribute (leaf or leaf-list)
    path: str = betterproto.string_field(1)
    # value_type is the data type of the attribute
    value_type: "_v2__.ValueType" = betterproto.enum_field(2)
    # units is the unit of measurement e.g. dB, mV
    units: str = betterproto.string_field(3)
    # description is an explaination of the meaning of the attribute
    description: str = betterproto.string_field(4)
    # mandatory shows whether the attribute is optional (false) or required
    # (true)
    mandatory: bool = betterproto.bool_field(5)
    # default is a default value used with optional attributes
    default: str = betterproto.string_field(6)
    # range is definition of the range of values a value is allowed
    range: List[str] = betterproto.string_field(7)
    # length is a defintion of the length restrictions for the attribute
    length: List[str] = betterproto.string_field(8)
    type_opts: List[int] = betterproto.uint64_field(9)
    is_a_key: bool = betterproto.bool_field(10)
    attr_name: str = betterproto.string_field(11)


@dataclass(eq=False, repr=False)
class ModelInfo(betterproto.Message):
    """ModelInfo is general information about a model plugin."""

    # name is the name given to the model plugin - no spaces and title case.
    name: str = betterproto.string_field(1)
    # version is the semantic version of the Plugin e.g. 1.0.0.
    version: str = betterproto.string_field(2)
    # model_data is a set of metadata about the YANG files that went in to
    # generating the model plugin. It includes name, version and organization for
    # each YANG file, similar to how they are represented in gNMI Capabilities.
    model_data: List["___gnmi__.ModelData"] = betterproto.message_field(3)
    # module is the name of the Model Plugin on the file system - usually ending
    # in .so.<version>.
    module: str = betterproto.string_field(4)
    # getStateMode is flag that defines how the "get state" operation works.  0)
    # means that no retrieval of state is attempted  1) means that the
    # synchronizer will make 2 requests to the device - one for      Get with
    # State and another for Get with Operational.  2) means that the synchronizer
    # will do a Get request comprising of each      one of the ReadOnlyPaths and
    # their sub paths. If there is a `list`      in any one of these paths it
    # will be sent down as is, expecting the      devices implementation of gNMI
    # will be able to expand wildcards.  3) means that the synchronizer will do a
    # Get request comprising of each      one of the ReadOnlyPaths and their sub
    # paths. If there is a `list`      in any one of these paths, a separate call
    # will be made first to find      all the instances in the list and a Get
    # including these expanded wildcards      will be sent down to the device.
    get_state_mode: int = betterproto.uint32_field(5)
    # read_only_path is all of the read only paths for the model plugin.
    read_only_path: List["ReadOnlyPath"] = betterproto.message_field(7)
    # read_write_path is all of the read write paths for the model plugin.
    read_write_path: List["ReadWritePath"] = betterproto.message_field(8)
    supported_encodings: List["___gnmi__.Encoding"] = betterproto.enum_field(9)


@dataclass(eq=False, repr=False)
class ModelPlugin(betterproto.Message):
    id: str = betterproto.string_field(1)
    endpoint: str = betterproto.string_field(2)
    info: "ModelInfo" = betterproto.message_field(3)
    status: str = betterproto.string_field(10)
    error: str = betterproto.string_field(11)


@dataclass(eq=False, repr=False)
class Chunk(betterproto.Message):
    """
    Chunk is for streaming a model plugin file to the server. There is a built
    in limit in gRPC of 4MB - plugin is usually around 20MB so break in to
    chunks of approx 1-2MB.
    """

    # so_file is the name being streamed.
    so_file: str = betterproto.string_field(1)
    # content is the bytes content.
    content: bytes = betterproto.bytes_field(2)


@dataclass(eq=False, repr=False)
class RegisterResponse(betterproto.Message):
    """RegisterResponse carries status of model plugin registration."""

    # name is name of the model plugin.
    name: str = betterproto.string_field(1)
    # version is the semantic version of the model plugin.
    version: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class ListModelsRequest(betterproto.Message):
    """
    ListModelsRequest carries data for querying registered model plugins.
    """

    # verbose option causes all of the ReadWrite and ReadOnly paths to be
    # included.
    verbose: bool = betterproto.bool_field(1)
    # An optional filter on the name of the model plugins to list.
    model_name: str = betterproto.string_field(2)
    # An optional filter on the version of the model plugins to list
    model_version: str = betterproto.string_field(3)


@dataclass(eq=False, repr=False)
class RollbackRequest(betterproto.Message):
    """
    RollbackRequest carries the index of the configuration change transaction
    to rollback.
    """

    # index of the transaction that should be rolled back
    index: int = betterproto.uint64_field(1)


@dataclass(eq=False, repr=False)
class RollbackResponse(betterproto.Message):
    """RollbackResponse carries the response of the rollback operation"""

    # ID of the rollback transaction
    id: str = betterproto.string_field(1)
    # index of the rollback transaction
    index: int = betterproto.uint64_field(2)


@dataclass(eq=False, repr=False)
class ListSnapshotsRequest(betterproto.Message):
    """
    ListSnapshotsRequest requests a list of snapshots for all devices and
    versions.
    """

    # subscribe indicates whether to subscribe to events (e.g. ADD, UPDATE, and
    # REMOVE) that occur after all devices have been streamed to the client
    subscribe: bool = betterproto.bool_field(1)
    # option to specify a specific device - if blank or '*' then select all Can
    # support `*` (match many chars) or '?' (match one char) as wildcard
    id: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class CompactChangesRequest(betterproto.Message):
    """
    CompactChangesRequest requests a compaction of the Network Change and
    Device Change stores
    """

    # retention_period is an optional duration of time counting back from the
    # present moment Network changes that were created during this period should
    # not be compacted Any network changes that are older should be compacted If
    # not specified the duration is 0
    retention_period: timedelta = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class CompactChangesResponse(betterproto.Message):
    """CompactChangesResponse is a response to the Compact Changes command"""

    pass


@dataclass(eq=False, repr=False)
class ModelInfoRequest(betterproto.Message):
    """ModelInfoRequest carries request for the model information"""

    pass


@dataclass(eq=False, repr=False)
class ModelInfoResponse(betterproto.Message):
    """ModelInfoResponse carries response for the model information query"""

    model_info: "ModelInfo" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class ValidateConfigRequest(betterproto.Message):
    """
    ValidateConfigRequest carries configuration data to be validated as a JSON
    blob
    """

    json: bytes = betterproto.bytes_field(1)


@dataclass(eq=False, repr=False)
class ValidateConfigResponse(betterproto.Message):
    """ValidateConfigResponse carries the result of the validation"""

    valid: bool = betterproto.bool_field(1)
    message: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class PathValuesRequest(betterproto.Message):
    """PathValuesRequest carries configuration change as a JSON blob"""

    path_prefix: str = betterproto.string_field(1)
    json: bytes = betterproto.bytes_field(2)


@dataclass(eq=False, repr=False)
class PathValuesResponse(betterproto.Message):
    """PathValuesResponse carries a list of typed path values"""

    path_values: List["_v2__.PathValue"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class GetTransactionRequest(betterproto.Message):
    id: str = betterproto.string_field(1)
    index: int = betterproto.uint64_field(2)


@dataclass(eq=False, repr=False)
class GetTransactionResponse(betterproto.Message):
    transaction: "_v2__.Transaction" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class ListTransactionsRequest(betterproto.Message):
    pass


@dataclass(eq=False, repr=False)
class ListTransactionsResponse(betterproto.Message):
    transaction: "_v2__.Transaction" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class WatchTransactionsRequest(betterproto.Message):
    id: str = betterproto.string_field(1)
    noreplay: bool = betterproto.bool_field(2)


@dataclass(eq=False, repr=False)
class WatchTransactionsResponse(betterproto.Message):
    event: "_v2__.TransactionEvent" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class GetConfigurationRequest(betterproto.Message):
    configuration_id: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class GetConfigurationResponse(betterproto.Message):
    configuration: "_v2__.Configuration" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class ListConfigurationsRequest(betterproto.Message):
    pass


@dataclass(eq=False, repr=False)
class ListConfigurationsResponse(betterproto.Message):
    configuration: "_v2__.Configuration" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class WatchConfigurationsRequest(betterproto.Message):
    configuration_id: str = betterproto.string_field(1)
    noreplay: bool = betterproto.bool_field(2)


@dataclass(eq=False, repr=False)
class WatchConfigurationsResponse(betterproto.Message):
    event: "_v2__.ConfigurationEvent" = betterproto.message_field(1)


class ConfigAdminServiceStub(betterproto.ServiceStub):
    async def upload_register_model(
        self, request_iterator: Union[AsyncIterable["Chunk"], Iterable["Chunk"]]
    ) -> "RegisterResponse":

        return await self._stream_unary(
            "/onos.config.admin.ConfigAdminService/UploadRegisterModel",
            request_iterator,
            Chunk,
            RegisterResponse,
        )

    async def list_registered_models(
        self, *, verbose: bool = False, model_name: str = "", model_version: str = ""
    ) -> AsyncIterator["ModelPlugin"]:

        request = ListModelsRequest()
        request.verbose = verbose
        request.model_name = model_name
        request.model_version = model_version

        async for response in self._unary_stream(
            "/onos.config.admin.ConfigAdminService/ListRegisteredModels",
            request,
            ModelPlugin,
        ):
            yield response

    async def rollback_transaction(self, *, index: int = 0) -> "RollbackResponse":

        request = RollbackRequest()
        request.index = index

        return await self._unary_unary(
            "/onos.config.admin.ConfigAdminService/RollbackTransaction",
            request,
            RollbackResponse,
        )

    async def list_snapshots(
        self, *, subscribe: bool = False, id: str = ""
    ) -> AsyncIterator["_snapshot_device__.Snapshot"]:

        request = ListSnapshotsRequest()
        request.subscribe = subscribe
        request.id = id

        async for response in self._unary_stream(
            "/onos.config.admin.ConfigAdminService/ListSnapshots",
            request,
            _snapshot_device__.Snapshot,
        ):
            yield response

    async def compact_changes(
        self, *, retention_period: timedelta = None
    ) -> "CompactChangesResponse":

        request = CompactChangesRequest()
        if retention_period is not None:
            request.retention_period = retention_period

        return await self._unary_unary(
            "/onos.config.admin.ConfigAdminService/CompactChanges",
            request,
            CompactChangesResponse,
        )


class ModelPluginServiceStub(betterproto.ServiceStub):
    async def get_model_info(self) -> "ModelInfoResponse":

        request = ModelInfoRequest()

        return await self._unary_unary(
            "/onos.config.admin.ModelPluginService/GetModelInfo",
            request,
            ModelInfoResponse,
        )

    async def validate_config(self, *, json: bytes = b"") -> "ValidateConfigResponse":

        request = ValidateConfigRequest()
        request.json = json

        return await self._unary_unary(
            "/onos.config.admin.ModelPluginService/ValidateConfig",
            request,
            ValidateConfigResponse,
        )

    async def get_path_values(
        self, *, path_prefix: str = "", json: bytes = b""
    ) -> "PathValuesResponse":

        request = PathValuesRequest()
        request.path_prefix = path_prefix
        request.json = json

        return await self._unary_unary(
            "/onos.config.admin.ModelPluginService/GetPathValues",
            request,
            PathValuesResponse,
        )


class TransactionServiceStub(betterproto.ServiceStub):
    async def get_transaction(
        self, *, id: str = "", index: int = 0
    ) -> "GetTransactionResponse":

        request = GetTransactionRequest()
        request.id = id
        request.index = index

        return await self._unary_unary(
            "/onos.config.admin.TransactionService/GetTransaction",
            request,
            GetTransactionResponse,
        )

    async def list_transactions(self) -> AsyncIterator["ListTransactionsResponse"]:

        request = ListTransactionsRequest()

        async for response in self._unary_stream(
            "/onos.config.admin.TransactionService/ListTransactions",
            request,
            ListTransactionsResponse,
        ):
            yield response

    async def watch_transactions(
        self, *, id: str = "", noreplay: bool = False
    ) -> AsyncIterator["WatchTransactionsResponse"]:

        request = WatchTransactionsRequest()
        request.id = id
        request.noreplay = noreplay

        async for response in self._unary_stream(
            "/onos.config.admin.TransactionService/WatchTransactions",
            request,
            WatchTransactionsResponse,
        ):
            yield response


class ConfigurationServiceStub(betterproto.ServiceStub):
    async def get_configuration(
        self, *, configuration_id: str = ""
    ) -> "GetConfigurationResponse":

        request = GetConfigurationRequest()
        request.configuration_id = configuration_id

        return await self._unary_unary(
            "/onos.config.admin.ConfigurationService/GetConfiguration",
            request,
            GetConfigurationResponse,
        )

    async def list_configurations(self) -> AsyncIterator["ListConfigurationsResponse"]:

        request = ListConfigurationsRequest()

        async for response in self._unary_stream(
            "/onos.config.admin.ConfigurationService/ListConfigurations",
            request,
            ListConfigurationsResponse,
        ):
            yield response

    async def watch_configurations(
        self, *, configuration_id: str = "", noreplay: bool = False
    ) -> AsyncIterator["WatchConfigurationsResponse"]:

        request = WatchConfigurationsRequest()
        request.configuration_id = configuration_id
        request.noreplay = noreplay

        async for response in self._unary_stream(
            "/onos.config.admin.ConfigurationService/WatchConfigurations",
            request,
            WatchConfigurationsResponse,
        ):
            yield response


class ConfigAdminServiceBase(ServiceBase):
    async def upload_register_model(
        self, request_iterator: AsyncIterator["Chunk"]
    ) -> "RegisterResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def list_registered_models(
        self, verbose: bool, model_name: str, model_version: str
    ) -> AsyncIterator["ModelPlugin"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def rollback_transaction(self, index: int) -> "RollbackResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def list_snapshots(
        self, subscribe: bool, id: str
    ) -> AsyncIterator["_snapshot_device__.Snapshot"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def compact_changes(
        self, retention_period: timedelta
    ) -> "CompactChangesResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_upload_register_model(self, stream: grpclib.server.Stream) -> None:
        request_kwargs = {"request_iterator": stream.__aiter__()}

        response = await self.upload_register_model(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_list_registered_models(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "verbose": request.verbose,
            "model_name": request.model_name,
            "model_version": request.model_version,
        }

        await self._call_rpc_handler_server_stream(
            self.list_registered_models,
            stream,
            request_kwargs,
        )

    async def __rpc_rollback_transaction(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "index": request.index,
        }

        response = await self.rollback_transaction(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_list_snapshots(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "subscribe": request.subscribe,
            "id": request.id,
        }

        await self._call_rpc_handler_server_stream(
            self.list_snapshots,
            stream,
            request_kwargs,
        )

    async def __rpc_compact_changes(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "retention_period": request.retention_period,
        }

        response = await self.compact_changes(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.config.admin.ConfigAdminService/UploadRegisterModel": grpclib.const.Handler(
                self.__rpc_upload_register_model,
                grpclib.const.Cardinality.STREAM_UNARY,
                Chunk,
                RegisterResponse,
            ),
            "/onos.config.admin.ConfigAdminService/ListRegisteredModels": grpclib.const.Handler(
                self.__rpc_list_registered_models,
                grpclib.const.Cardinality.UNARY_STREAM,
                ListModelsRequest,
                ModelPlugin,
            ),
            "/onos.config.admin.ConfigAdminService/RollbackTransaction": grpclib.const.Handler(
                self.__rpc_rollback_transaction,
                grpclib.const.Cardinality.UNARY_UNARY,
                RollbackRequest,
                RollbackResponse,
            ),
            "/onos.config.admin.ConfigAdminService/ListSnapshots": grpclib.const.Handler(
                self.__rpc_list_snapshots,
                grpclib.const.Cardinality.UNARY_STREAM,
                ListSnapshotsRequest,
                _snapshot_device__.Snapshot,
            ),
            "/onos.config.admin.ConfigAdminService/CompactChanges": grpclib.const.Handler(
                self.__rpc_compact_changes,
                grpclib.const.Cardinality.UNARY_UNARY,
                CompactChangesRequest,
                CompactChangesResponse,
            ),
        }


class ModelPluginServiceBase(ServiceBase):
    async def get_model_info(self) -> "ModelInfoResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def validate_config(self, json: bytes) -> "ValidateConfigResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def get_path_values(
        self, path_prefix: str, json: bytes
    ) -> "PathValuesResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_get_model_info(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        response = await self.get_model_info(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_validate_config(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "json": request.json,
        }

        response = await self.validate_config(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_get_path_values(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "path_prefix": request.path_prefix,
            "json": request.json,
        }

        response = await self.get_path_values(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.config.admin.ModelPluginService/GetModelInfo": grpclib.const.Handler(
                self.__rpc_get_model_info,
                grpclib.const.Cardinality.UNARY_UNARY,
                ModelInfoRequest,
                ModelInfoResponse,
            ),
            "/onos.config.admin.ModelPluginService/ValidateConfig": grpclib.const.Handler(
                self.__rpc_validate_config,
                grpclib.const.Cardinality.UNARY_UNARY,
                ValidateConfigRequest,
                ValidateConfigResponse,
            ),
            "/onos.config.admin.ModelPluginService/GetPathValues": grpclib.const.Handler(
                self.__rpc_get_path_values,
                grpclib.const.Cardinality.UNARY_UNARY,
                PathValuesRequest,
                PathValuesResponse,
            ),
        }


class TransactionServiceBase(ServiceBase):
    async def get_transaction(self, id: str, index: int) -> "GetTransactionResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def list_transactions(self) -> AsyncIterator["ListTransactionsResponse"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def watch_transactions(
        self, id: str, noreplay: bool
    ) -> AsyncIterator["WatchTransactionsResponse"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_get_transaction(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "id": request.id,
            "index": request.index,
        }

        response = await self.get_transaction(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_list_transactions(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        await self._call_rpc_handler_server_stream(
            self.list_transactions,
            stream,
            request_kwargs,
        )

    async def __rpc_watch_transactions(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "id": request.id,
            "noreplay": request.noreplay,
        }

        await self._call_rpc_handler_server_stream(
            self.watch_transactions,
            stream,
            request_kwargs,
        )

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.config.admin.TransactionService/GetTransaction": grpclib.const.Handler(
                self.__rpc_get_transaction,
                grpclib.const.Cardinality.UNARY_UNARY,
                GetTransactionRequest,
                GetTransactionResponse,
            ),
            "/onos.config.admin.TransactionService/ListTransactions": grpclib.const.Handler(
                self.__rpc_list_transactions,
                grpclib.const.Cardinality.UNARY_STREAM,
                ListTransactionsRequest,
                ListTransactionsResponse,
            ),
            "/onos.config.admin.TransactionService/WatchTransactions": grpclib.const.Handler(
                self.__rpc_watch_transactions,
                grpclib.const.Cardinality.UNARY_STREAM,
                WatchTransactionsRequest,
                WatchTransactionsResponse,
            ),
        }


class ConfigurationServiceBase(ServiceBase):
    async def get_configuration(
        self, configuration_id: str
    ) -> "GetConfigurationResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def list_configurations(self) -> AsyncIterator["ListConfigurationsResponse"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def watch_configurations(
        self, configuration_id: str, noreplay: bool
    ) -> AsyncIterator["WatchConfigurationsResponse"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_get_configuration(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "configuration_id": request.configuration_id,
        }

        response = await self.get_configuration(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_list_configurations(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        await self._call_rpc_handler_server_stream(
            self.list_configurations,
            stream,
            request_kwargs,
        )

    async def __rpc_watch_configurations(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "configuration_id": request.configuration_id,
            "noreplay": request.noreplay,
        }

        await self._call_rpc_handler_server_stream(
            self.watch_configurations,
            stream,
            request_kwargs,
        )

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.config.admin.ConfigurationService/GetConfiguration": grpclib.const.Handler(
                self.__rpc_get_configuration,
                grpclib.const.Cardinality.UNARY_UNARY,
                GetConfigurationRequest,
                GetConfigurationResponse,
            ),
            "/onos.config.admin.ConfigurationService/ListConfigurations": grpclib.const.Handler(
                self.__rpc_list_configurations,
                grpclib.const.Cardinality.UNARY_STREAM,
                ListConfigurationsRequest,
                ListConfigurationsResponse,
            ),
            "/onos.config.admin.ConfigurationService/WatchConfigurations": grpclib.const.Handler(
                self.__rpc_watch_configurations,
                grpclib.const.Cardinality.UNARY_STREAM,
                WatchConfigurationsRequest,
                WatchConfigurationsResponse,
            ),
        }


from .. import v2 as _v2__
from .... import gnmi as ___gnmi__
from ..snapshot import device as _snapshot_device__
