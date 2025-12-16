# conf1
**PERFORMANCE EVALUATION OF MULTI-PAXOS AND MENCIUS ALGORITHMS IN DISTRIBUTED SYSTEMS**
* Author: Vipul Kumar Bondugula
* Published In : Intern************
* Publication Date: 05-2025
* E-ISSN: 2****
* Impact Factor: ****
* Link:

**Abstract:**\
This paper addresses performance and scalability limitations in distributed consensus systems caused by the leader-centric design of the Multi-Paxos algorithm. It examines how reliance on a single leader for proposal coordination creates throughput bottlenecks, increases recovery delays during leader failures, and limits system scalability as the number of nodes grows. The study emphasizes the trade-offs of Multi-Paxos, including strong consistency and fault tolerance at the cost of reduced throughput and increased leader workload. By contrasting Multi-Paxos with the Mencius algorithm, which enables decentralized proposal handling and concurrent decision-making, the proposed analysis highlights significant improvements in throughput and load distribution. The paper underscores the need for decentralized consensus mechanisms to achieve higher performance and scalability in large-scale distributed systems while maintaining consistency guarantees.

**Key Contributions:**
* **Consensus Bottleneck Analysis:**\
Examined the throughput limitations of leader-based consensus caused by centralized coordination and increasing leader workload in Multi-Paxos.

* **Decentralized Proposal Strategy:**\
Analyzed the Mencius consensus approach, where multiple nodes propose concurrently, reducing reliance on a single coordinator and improving parallelism.
  
* **Comparative Throughput Evaluation:** \
  Conducted a detailed performance comparison of Multi-Paxos and Mencius across varying cluster sizes, demonstrating consistently higher throughput for Mencius.
  
* **Design and Experimental Validation:**\
  Designed, implemented, and validated simulation-based experiments to measure throughput behavior and scalability characteristics of both consensus algorithms.

**Relevance & Real-World Impact**
* **Improved System Throughput:**\
Achieved significantly higher operations per second by eliminating leader bottlenecks through decentralized consensus.

* **Better Scalability for Large Clusters:**\
Enabled more efficient scaling of distributed systems as cluster size increases, with gradual performance degradation instead of sharp drops.

* **Reduced Coordination Overhead:** \
    Distributed proposal responsibilities across nodes, minimizing idle time and improving overall resource utilization.
  
* **Academic and Practical Value:** \
    Provides clear experimental evidence and implementation insights useful for research, teaching, and system design involving distributed consensus algorithms.
  
**Experimental Results (Summary)**:

  | Nodes | Multi-Paxos (k ops/sec) | Mencius (k ops/sec) | Reduction (%)   |
  |-------|-------------------------| --------------------| ----------------|
  | 3     |  60                     | 110                 | 83.33           |
  | 5     |  75                     | 135                 | 80.00           |
  | 7     |  65                     | 123                 | 89.23           |
  | 9     |  57                     | 105                 | 84.21           |
  | 11    |  52                     | 95                  | 82.69           |

**Citation** \
PERFORMANCE EVALUATION OF MULTI-PAXOS AND MENCIUS ALGORITHMS IN DISTRIBUTED SYSTEMS
* Vipul Kumar Bondugula
* Inter***********echnology 
* E-ISSN 2*******
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.*******/ \
**Author Contact** \
**LinkedIn**: https://www.linkedin.com/in/vipul-b-18468a19/ | **Email**: vipulreddy574@gmail.com






