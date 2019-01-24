InnoDB多版本(MVCC)实现简要分析
http://hedengcheng.com/?p=148

MTR:
mtr是最小的事物单位,保证一个物理事物的完成。

所有的物理操作,包括undoLog的redoLog都需要其来完成


之前的疑问:
mtr 在 space management 中担任的角色其实是无关事物的
mtr 修改了一个链表 此时提redo buffer 此时修改是可读的因为redoLog是连续的
mtr_commit之后,后续的对space的操作才能在此之后刷到redoLog buffer,如果第一个链表
mtr操作没有提交那么内存就不变,如果已经提交,修改即生效。如果 redo buffer没有刷盘也不要紧
因为后续因为链表改变所产生的数据变化都是未刷盘的。

* 1 一个mtr产生的redoLog是连续的
* 2 一个事物可能包含多个mtr 事物的状态由undolog事物状态决定
* 3 redoLog的刷新是事物无关的 log可能有完成的T1也有可能有未完成的T2
