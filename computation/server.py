from concurrent import futures
import logging

import grpc

import math_pb2
import math_pb2_grpc


class Calculator(math_pb2_grpc.CalculatorServicer):

    def Sum(self, request, context):
        print('Start Sum computation...')
        result = request.operand1 + request.operand2
        for i in range(10):
            if i == 9:
                yield math_pb2.SumResponse(finished=True, result=result)
                print('Finished Sum computation, result={}'.format(result))
            else:
                yield math_pb2.SumResponse(finished=False, result=-1)
                print('Still calculating Sum...')


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
