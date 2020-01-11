# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import math_pb2 as math__pb2


class ArithmeticStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Sum = channel.unary_unary(
        '/math.Arithmetic/Sum',
        request_serializer=math__pb2.SumRequest.SerializeToString,
        response_deserializer=math__pb2.SumResponse.FromString,
        )


class ArithmeticServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Sum(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_ArithmeticServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Sum': grpc.unary_unary_rpc_method_handler(
          servicer.Sum,
          request_deserializer=math__pb2.SumRequest.FromString,
          response_serializer=math__pb2.SumResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'math.Arithmetic', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))