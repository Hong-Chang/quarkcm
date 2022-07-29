/*
Copyright 2022 quarkcm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// use crate::constants::*;
// use crate::rdma_ctrlconn::*;
use crate::RDMA_CTLINFO;
// use svc_client::quark_cm_service_client::QuarkCmServiceClient;
// use svc_client::MaxResourceVersionMessage;
// use svc_client::PodMessage;
use tokio::time::*;
// use tonic::Request;

pub mod svc_client {
    tonic::include_proto!("quarkcmsvc");
}

#[derive(Debug)]
pub struct PodInformer {
    pub max_resource_version: i32,
}

impl PodInformer {
    pub async fn new() -> Result<Self, Box<dyn std::error::Error>> {
        let informer = Self {
            max_resource_version: 0,
        };
        // let mut client = QuarkCmServiceClient::connect(GRPC_SERVER_ADDRESS).await?;

        // let ref pods_message = client.list_pod(()).await?.into_inner().pods;
        // if pods_message.len() > 0 {
        //     for pod_message in pods_message {
        //         informer.handle(pod_message);
        //     }
        // }

        Ok(informer)
    }
}

impl PodInformer {
    pub async fn run(&mut self) -> Result<(), Box<dyn std::error::Error>> {
        loop {
            match self.run_watch().await {
                Ok(_) => {}
                Err(e) => {
                    println!("Pod watch error: {:?}", e);
                }
            }

            if *RDMA_CTLINFO.exiting.lock() {
                println!("PodInformer exit!");
                break;
            } else {
                // println!("PodInformer sleeps 1 second for next watch session.");
                sleep(Duration::from_secs(1)).await;
            }
        }
        Ok(())
    }

    async fn run_watch(&mut self) -> Result<(), Box<dyn std::error::Error>> {
        // let mut client = QuarkCmServiceClient::connect(GRPC_SERVER_ADDRESS).await?;
        // let mut pod_stream = client
        //     .watch_pod(Request::new(MaxResourceVersionMessage {
        //         max_resource_version: self.max_resource_version,
        //         // max_resource_version: 0,
        //     }))
        //     .await?
        //     .into_inner();

        // while let Some(pod_message) = pod_stream.message().await? {
        //     self.handle(&pod_message);
        // }
        Ok(())
    }

    // fn handle(&mut self, pod_message: &PodMessage) {
    //     let ip = pod_message.ip;
    //     let mut pods_map = RDMA_CTLINFO.pods.lock();
    //     if pod_message.event_type == EVENT_TYPE_SET {
    //         let pod = Pod {
    //             key: pod_message.key.clone(),
    //             ip: ip,
    //             node_name: pod_message.node_name.clone(),
    //             resource_version: pod_message.resource_version,
    //         };
    //         pods_map.insert(ip, pod);
    //         if pod_message.resource_version > self.max_resource_version {
    //             self.max_resource_version = pod_message.resource_version;
    //         }
    //     } else if pod_message.event_type == EVENT_TYPE_DELETE {
    //         if pods_map.contains_key(&ip) {
    //             if pods_map[&ip].resource_version < pod_message.resource_version {
    //                 pods_map.remove(&ip);
    //             }
    //         }
    //     }
    //     if pod_message.resource_version > self.max_resource_version {
    //         self.max_resource_version = pod_message.resource_version;
    //     }
    //     println!("Handled Pod: {:?}", pod_message);
    //     println!("Debug: pods_map len:{} {:?}", pods_map.len(), pods_map);
    // }
}
