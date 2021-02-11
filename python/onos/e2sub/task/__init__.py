# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/e2sub/task/task.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import AsyncIterator, List, Optional

import betterproto
import grpclib


class Phase(betterproto.Enum):
    """Phase is a subscription task phase"""

    # OPEN is a subscription task open phase
    OPEN = 0
    # CLOSE is a subscription task close phase
    CLOSE = 1


class Status(betterproto.Enum):
    """Status is a subscription task status"""

    # PENDING indicates the subscription task phase is pending
    PENDING = 0
    # COMPLETE indicates the subscription task phase is complete
    COMPLETE = 1
    # FAILED indicates the subscription task phase failed
    FAILED = 2


class Cause(betterproto.Enum):
    """Cause is a failure cause"""

    CAUSE_UNKNOWN = 0
    CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD = 1
    CAUSE_MISC_HARDWARE_FAILURE = 2
    CAUSE_MISC_OM_INTERVENTION = 3
    CAUSE_MISC_UNSPECIFIED = 4
    CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR = 5
    CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT = 6
    CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY = 7
    CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE = 8
    CAUSE_PROTOCOL_SEMANTIC_ERROR = 9
    CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE = 10
    CAUSE_PROTOCOL_UNSPECIFIED = 11
    CAUSE_RIC_RAN_FUNCTION_ID_INVALID = 12
    CAUSE_RIC_ACTION_NOT_SUPPORTED = 13
    CAUSE_RIC_EXCESSIVE_ACTIONS = 14
    CAUSE_RIC_DUPLICATE_ACTION = 15
    CAUSE_RIC_DUPLICATE_EVENT = 16
    CAUSE_RIC_FUNCTION_RESOURCE_LIMIT = 17
    CAUSE_RIC_REQUEST_ID_UNKNOWN = 18
    CAUSE_RIC_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE = 19
    CAUSE_RIC_CONTROL_MESSAGE_INVALID = 20
    CAUSE_RIC_CALL_PROCESS_ID_INVALID = 21
    CAUSE_RIC_UNSPECIFIED = 22
    CAUSE_RICSERVICE_FUNCTION_NOT_REQUIRED = 23
    CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS = 24
    CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT = 25
    CAUSE_TRANSPORT_UNSPECIFIED = 26
    CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE = 27


class EventType(betterproto.Enum):
    """Type of change"""

    NONE = 0
    CREATED = 1
    UPDATED = 2
    REMOVED = 3


@dataclass(eq=False, repr=False)
class Lifecycle(betterproto.Message):
    """Lifecycle is a subscription task status"""

    phase: "Phase" = betterproto.enum_field(1)
    status: "Status" = betterproto.enum_field(2)
    failure: "Failure" = betterproto.message_field(3)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Failure(betterproto.Message):
    """Failure is a subscription failure status"""

    cause: "Cause" = betterproto.enum_field(1)
    message: str = betterproto.string_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class SubscriptionTask(betterproto.Message):
    """
    SubscriptionTask is a task representing a subscription between an E2
    termination and an E2 node
    """

    id: str = betterproto.string_field(1)
    revision: int = betterproto.uint64_field(2)
    subscription_id: str = betterproto.string_field(3)
    endpoint_id: str = betterproto.string_field(4)
    lifecycle: "Lifecycle" = betterproto.message_field(5)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Event(betterproto.Message):
    """Event is a SubscriptionTask event"""

    type: "EventType" = betterproto.enum_field(1)
    task: "SubscriptionTask" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetSubscriptionTaskRequest(betterproto.Message):
    """
    GetSubscriptionTaskRequest is a request for getting existing
    SubscriptionTask
    """

    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetSubscriptionTaskResponse(betterproto.Message):
    """
    GetSubscriptionTaskResponse is a response with invormation about a
    requested SubscriptionTask
    """

    task: "SubscriptionTask" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ListSubscriptionTasksRequest(betterproto.Message):
    """
    ListSubscriptionTasksRequest is a request to list all available
    SubscriptionTasks
    """

    subscription_id: str = betterproto.string_field(1)
    endpoint_id: str = betterproto.string_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ListSubscriptionTasksResponse(betterproto.Message):
    """
    ListSubscriptionTasksResponse is a response to list all available
    SubscriptionTasks
    """

    tasks: List["SubscriptionTask"] = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class WatchSubscriptionTasksRequest(betterproto.Message):
    """
    WatchSubscriptionTasksRequest is a request to receive a stream of all
    SubscriptionTask changes.
    """

    noreplay: bool = betterproto.bool_field(1)
    subscription_id: str = betterproto.string_field(2)
    endpoint_id: str = betterproto.string_field(3)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class WatchSubscriptionTasksResponse(betterproto.Message):
    """
    WatchSubscriptionTasksResponse is a response indicating a change in the
    available SubscriptionTasks.
    """

    event: "Event" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class UpdateSubscriptionTaskRequest(betterproto.Message):
    """
    UpdateSubscriptionTaskRequest is a request for updating a SubscriptionTask
    status
    """

    task: "SubscriptionTask" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class UpdateSubscriptionTaskResponse(betterproto.Message):
    """
    UpdateSubscriptionTaskResponse is a response to updating a SubscriptionTask
    status
    """

    pass

    def __post_init__(self) -> None:
        super().__post_init__()


class E2SubscriptionTaskServiceStub(betterproto.ServiceStub):
    """
    E2SubscriptionTaskService manages subscription tasks between E2 termination
    points and E2 nodes
    """

    async def get_subscription_task(
        self, *, id: str = ""
    ) -> "GetSubscriptionTaskResponse":
        """GetSubscriptionTask retrieves information about a specific task"""

        request = GetSubscriptionTaskRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.e2sub.task.E2SubscriptionTaskService/GetSubscriptionTask",
            request,
            GetSubscriptionTaskResponse,
        )

    async def list_subscription_tasks(
        self, *, subscription_id: str = "", endpoint_id: str = ""
    ) -> "ListSubscriptionTasksResponse":
        """
        ListSubscriptionTasks returns the list of currently registered E2
        Tasks.
        """

        request = ListSubscriptionTasksRequest()
        request.subscription_id = subscription_id
        request.endpoint_id = endpoint_id

        return await self._unary_unary(
            "/onos.e2sub.task.E2SubscriptionTaskService/ListSubscriptionTasks",
            request,
            ListSubscriptionTasksResponse,
        )

    async def watch_subscription_tasks(
        self,
        *,
        noreplay: bool = False,
        subscription_id: str = "",
        endpoint_id: str = "",
    ) -> AsyncIterator["WatchSubscriptionTasksResponse"]:
        """
        WatchSubscriptionTasks returns a stream of changes in the set of
        available E2 Tasks.
        """

        request = WatchSubscriptionTasksRequest()
        request.noreplay = noreplay
        request.subscription_id = subscription_id
        request.endpoint_id = endpoint_id

        async for response in self._unary_stream(
            "/onos.e2sub.task.E2SubscriptionTaskService/WatchSubscriptionTasks",
            request,
            WatchSubscriptionTasksResponse,
        ):
            yield response

    async def update_subscription_task(
        self, *, task: "SubscriptionTask" = None
    ) -> "UpdateSubscriptionTaskResponse":
        """UpdateSubscriptionTask updates a task status"""

        request = UpdateSubscriptionTaskRequest()
        if task is not None:
            request.task = task

        return await self._unary_unary(
            "/onos.e2sub.task.E2SubscriptionTaskService/UpdateSubscriptionTask",
            request,
            UpdateSubscriptionTaskResponse,
        )
