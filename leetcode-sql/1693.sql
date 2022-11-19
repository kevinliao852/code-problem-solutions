-- Write your MySQL query statement below
SELECT
  date_id,
  make_name,
  COUNT(DISTINCT lead_id) unique_leads,
  COUNT(DISTINCT partner_id) unique_partners
FROM
  DailySales
GROUP BY
  date_id,
  make_name
