# conf1
**PERFORMANCE EVALUATION OF MULTI-PAXOS AND MENCIUS ALGORITHMS IN DISTRIBUTED SYSTEMS**
* Author: Vipul Reddy
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

  | Nodes | Multi-Version Concurrency Control Storage | Optimistic Concurrency Control | Reduction (%)   |
  |-------|-------------------------------------------| -------------------------------| ----------------|
  | 3     |  3                                        | 1                              | 66.67           |
  | 5     |  5                                        | 1.5                            | 70.00           |
  | 7     |  7                                        | 2                              | 71.43           |
  | 9     |  9                                        | 2.5                            | 72.22           |
  | 11    |  11                                       | 3                              | 72.73           |

**Citation** \
STORAGE OPTIMIZATION IN DISTRIBUTED ENVIRONMENTS USING OPTIMISTIC CONCURRENCY CONTROL
* Vipul R 
* International Journal on Science and Technology 
* E-ISSN 2229-7677
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijsat.org/ \
**Author Contact** \
**LinkedIn**: http://linkedin.com/in/Please add here | **Email**: please keep email id @gmail.com






