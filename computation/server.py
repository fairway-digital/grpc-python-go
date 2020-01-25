from concurrent import futures
import logging

import grpc

import time

import math_pb2
import math_pb2_grpc


class Calculator(math_pb2_grpc.CalculatorServicer):

    def Sum(self, request, context):
        print('Start Sum computation...')
        result = request.operand1 + request.operand2
        for i in range(5):
            if i == 4:
                yield math_pb2.SumResponse(finished=True, result=result)
                print('Finished Sum computation, result={}'.format(result))
            else:
                yield math_pb2.SumResponse(finished=False, result=-1)
                print('Still calculating Sum...')
            time.sleep(1)


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
