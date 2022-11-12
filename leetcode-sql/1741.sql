-- Write your MySQL query statement below
WITH
  cte1 AS (
    SELECT
      event_day day,
      emp_id,
      (out_time - in_time) total_time
    FROM
      Employees
  )
SELECT
  day,
  emp_id,
  SUM(total_time) total_time
FROM
  cte1
GROUP BY
  day,
  emp_id
