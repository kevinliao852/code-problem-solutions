-- Write your MySQL query statement below
WITH
  cte1 AS (
    SELECT
      account,
      SUM(amount) balance
    FROM
      Transactions
    GROUP BY
      account
  )
SELECT
  name NAME,
  balance BALANCE
FROM
  Users
  INNER JOIN cte1 ON cte1.account = Users.account
WHERE
  balance > 10000
