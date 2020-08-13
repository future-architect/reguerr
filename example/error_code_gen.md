# Error Code List

+------+--------------------------+-----------+------------+-----------------------------+
| CODE |           NAME           | LOGLEVEL  | STATUSCODE |           FORMAT            |
+------+--------------------------+-----------+------------+-----------------------------+
| 1001 | PermissionDeniedErr      | default   | default    | permission denied           |
| 1002 | UpdateConflictErr        | default   | default    | other user updated: key=%s  |
| 1003 | InvalidInputParameterErr | default   | default    | invalid input parameter: %v |
| 1005 | UserTypeUnregisterErr    | default   | default    | not found user type         |
| 1004 | NotFoundOperationIDErr   | WarnLevel |        404 | not found operation id      |
+------+--------------------------+-----------+------------+-----------------------------+
~~~~