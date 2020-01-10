from __future__ import print_function
import logging

import grpc

import sum_pb2
import sum_pb2_grpc


def run():
    with grpc.insecure_channel(
            target='localhost:50051',
            options=[('grpc.lb_policy_name', 'pick_first'),
                     ('grpc.enable_retries', 0), ('grpc.keepalive_timeout_ms',
                                                  10000)]) as channel:
        stub = sum_pb2_grpc.ArithmeticStub(channel)

        response = stub.Sum(sum_pb2.SumRequest(
            operand1=1,
            operand2=1
        ), timeout=10)
    print("Arithmetic sum client received 1 + 1 = " + str(response.result))


if __name__ == '__main__':
    logging.basicConfig()
    run()
