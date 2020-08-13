# Error Code List

+------+--------------------------+-----------+------------+-----------------------------+
| CODE |           NAME           | LOGLEVEL  | STATUSCODE |           FORMAT            |
+------+--------------------------+-----------+------------+-----------------------------+
| 1001 | PermissionDeniedErr      | Level(0)  |          0 | permission denied           |
| 1002 | UpdateConflictErr        | Level(0)  |          0 | other user updated: key=%s  |
| 1003 | InvalidInputParameterErr | Level(0)  |          0 | invalid input parameter: %v |
| 1005 | UserTypeUnregisterErr    | Level(0)  |          0 | not found user type         |
| 1004 | NotFoundOperationIDErr   | WarnLevel |        404 | not found operation id      |
+------+--------------------------+-----------+------------+-----------------------------+
