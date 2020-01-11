from concurrent import futures
import logging

import grpc

import math_pb2
import math_pb2_grpc


class Calculator(math_pb2_grpc.CalculatorServicer):

    def Sum(self, request, context):
        result = request.operand1 + request.operand2
        return math_pb2.SumResponse(result=result)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    math_pb2_grpc.add_CalculatorServicer_to_server(Calculator(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print('Server started, wait for calculation request')
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
