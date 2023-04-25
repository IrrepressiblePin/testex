## Подзадача 2
Статистика по дням. Посчитать: число событий по дням, число показов, число кликов, число уникальных объявлений, уникальных кампаний.
``` sql
SELECT 
    date, 
    countIf(event = 'click') AS clicks_count, 
    countIf(event = 'view') AS views_count, 
    countIf(ad_cost_type='CPC') AS unique_ad_cost_types_cpc, 
    countIf(ad_cost_type='CPM') AS unique_ad_cost_types_cpm, 
    COUNT(DISTINCT compain_union_id) AS unique_compain_union_ids 
FROM 
    table 
GROUP BY 
    date;
```

Найти объявления по которым показ произошел после клика
``` sql
SELECT 
    *
FROM 
    table AS a1 
JOIN 
    adv_data AS a2 ON a1.compain_union_id = a2.compain_union_id 
    AND a1.client_union_id = a2.client_union_id
    AND a1.ad_id = a2.ad_id
    AND a1.platform = a2.platform
WHERE 
    a1.event = 'view' AND a2.event = 'click' AND a1.time > a2.time

```

## Подзадача 3

Найти для каждого сотрудника сумму последней полученной зарплаты (13% от ставки(employee.gross_salary) + salary.bonus). Не исключать сотрудника из выборки, если он еще ни разу не получал ЗП. Бонусы выплачиваются вместе с основной частью, поэтому если в таблице salary есть бонус (даже нулевой) это говорит о том, что сотрудник хотя бы один раз получал ЗП. В ответе вывести отдел, id и имя сотрудника, ЗП (формула выше) и дату (только последнюю). Подсказка: стараться не использовать вложенные запросы, если это возможно. использовать вложенный
``` sql
SELECT e.department, e.id, e.name, 
       e.gross_salary * (1-0.13) + s.bonus AS salary, 
       s.date
FROM employee e 
LEFT JOIN (
  SELECT employee_id, bonus, date
  FROM salary s1
  WHERE date = (
    SELECT MAX(date)
    FROM salary s2
    WHERE s1.employee_id = s2.employee_id
  )
) s ON e.id = s.employee_id
ORDER BY 
  e.id;
```
