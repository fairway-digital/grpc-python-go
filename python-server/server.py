from concurrent import futures
import logging

import grpc

import sum_pb2
import sum_pb2_grpc


class Arithmetic(sum_pb2_grpc.ArithmeticServicer):

    def Sum(self, request, context):
        result = request.operand1 + request.operand2
        return sum_pb2.SumResponse(result=result)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    sum_pb2_grpc.add_ArithmeticServicer_to_server(Arithmetic(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
