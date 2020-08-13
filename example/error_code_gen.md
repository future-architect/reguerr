# Error Code List

+------+--------------------------+----------+------------+-----------------------------+
| CODE |           NAME           | LOGLEVEL | STATUSCODE |           FORMAT            |
+------+--------------------------+----------+------------+-----------------------------+
| 1001 | PermissionDeniedErr      |        0 |          0 | permission denied           |
| 1002 | UpdateConflictErr        |        0 |          0 | other user updated: key=%s  |
| 1003 | InvalidInputParameterErr |        0 |          0 | invalid input parameter: %v |
| 1005 | UserTypeUnregisterErr    |        0 |          0 | not found user type         |
| 1004 | NotFoundOperationIDErr   |        4 |        404 | not found operation id      |
+------+--------------------------+----------+------------+-----------------------------+
