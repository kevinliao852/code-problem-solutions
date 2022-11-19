-- Write your MySQL query statement below
WITH
  cte1 AS (
    SELECT
      visit_id no_trans_visit_id
    FROM
      Transactions
    GROUP BY
      visit_id
  ),
  cte2 AS (
    SELECT
      visit_id,
      customer_id
    FROM
      Visits
      LEFT JOIN cte1 ON cte1.no_trans_visit_id = visit_id
    WHERE
      no_trans_visit_id IS NULL
  ) (
    SELECT
      customer_id,
      COUNT(customer_id) count_no_trans
    FROM
      cte2
    GROUP BY
      customer_id
  )
