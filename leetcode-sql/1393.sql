-- Write your MySQL query statement below
WITH
  cte1 AS (
    SELECT
      stock_name,
      operation,
      SUM(
        CASE
          WHEN operation = 'Sell' THEN price
          ELSE price * -1
        END
      ) capital_gain_loss
    FROM
      Stocks
    GROUP BY
      stock_name,
      operation
  )
SELECT
  stock_name,
  SUM(capital_gain_loss) capital_gain_loss
FROM
  cte1
GROUP BY
  stock_name
