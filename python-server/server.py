from concurrent import futures
import logging

import grpc

from protos import math_pb2
from protos import math_pb2_grpc


class Arithmetic(math_pb2_grpc.ArithmeticServicer):

    def Sum(self, request, context):
        result = request.operand1 + request.operand2
        return math_pb2.SumResponse(result=result)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    math_pb2_grpc.add_ArithmeticServicer_to_server(Arithmetic(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
