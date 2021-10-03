# coding: utf-8

"""
    tensorflow

    Python SDK for tensorflow  # noqa: E501

    The version of the OpenAPI document: v1.3.0
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

from kubeflow.training.models import *
from kubeflow.training.models.v1_xg_boost_job import V1XGBoostJob  # noqa: E501
from kubeflow.training.rest import ApiException

class TestV1XGBoostJob(unittest.TestCase):
    """V1XGBoostJob unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V1XGBoostJob
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = kubeflow.training.models.v1_xg_boost_job.V1XGBoostJob()  # noqa: E501
        if include_optional :
            return V1XGBoostJob(
                api_version = '0', 
                kind = '0', 
                metadata = None, 
                spec = V1XGBoostJobSpec(
                    run_policy = V1RunPolicy(
                        active_deadline_seconds = 56, 
                        backoff_limit = 56, 
                        clean_pod_policy = '0', 
                        scheduling_policy = V1SchedulingPolicy(
                            min_available = 56, 
                            min_resources = {
                                'key' : None
                                }, 
                            priority_class = '0', 
                            queue = '0', ), 
                        ttl_seconds_after_finished = 56, ), 
                    xgb_replica_specs = {
                        'key' : V1ReplicaSpec(
                            replicas = 56, 
                            restart_policy = '0', 
                            template = None, )
                        }, ), 
                status = V1JobStatus(
                    completion_time = None, 
                    conditions = [
                        V1JobCondition(
                            last_transition_time = None, 
                            last_update_time = None, 
                            message = '0', 
                            reason = '0', 
                            status = '0', 
                            type = '0', )
                        ], 
                    last_reconcile_time = None, 
                    replica_statuses = {
                        'key' : V1ReplicaStatus(
                            active = 56, 
                            failed = 56, 
                            succeeded = 56, )
                        }, 
                    start_time = None, )
            )
        else :
            return V1XGBoostJob(
        )

    def testV1XGBoostJob(self):
        """Test V1XGBoostJob"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
