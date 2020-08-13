# Error Code List

+------+--------------------------+----------+------------+-----------------------------+
| CODE |           NAME           | LOGLEVEL | STATUSCODE |           FORMAT            |
+------+--------------------------+----------+------------+-----------------------------+
| 1001 | PermissionDeniedErr      | Error    |        500 | permission denied           |
| 1002 | UpdateConflictErr        | Error    |        500 | other user updated: key=%s  |
| 1003 | InvalidInputParameterErr | Error    |        500 | invalid input parameter: %v |
| 1005 | UserTypeUnregisterErr    | Error    |        500 | not found user type         |
| 1004 | NotFoundOperationIDErr   | Warn     |        404 | not found operation id      |
+------+--------------------------+----------+------------+-----------------------------+
