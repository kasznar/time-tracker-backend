# time-tracker-backend


```shell
 docker run --name gorm-mysql -e MYSQL_ROOT_PASSWORD=12345678 -p 3306:3306 -d mysql
```

```shell
CREATE DATABASE GO_DB;
```

dto
```
[
    {
        date: '2021-03-12'
        team: [
            {
                userId: 12
                workday: {
                    type: 'workday'
                    summary: 'dolgoztam soksat nagyon'
                }
            },
            {
                userId: 22
                workday: null
            }

        ]
    }
]
```


query result
```

date, userId, workday_type, worday_summary

& group by date


```


```mysql
select calendar_days.date, users.id
    from users, calendar_days
    where calendar_days.date
    between DATE('2021-07-01') and  DATE('2021-07-31')
    order by calendar_days.date
```

```mysql
select calendar_days.date, users.id, users.name, w.summary, wt.name
from calendar_days
    cross join users
    left join workdays w on calendar_days.id = w.calendar_day_id and users.id = w.user_id
    left join workday_types wt on w.workday_type_id = wt.id
    left join user_teams ut on users.id = ut.user_id
where
      ut.team_id = 1 and
      calendar_days.date between DATE('2021-01-01') and  DATE('2021-01-31')
```


```shell
 export TIME_TRACKER_DB_USER=root
 export TIME_TRACKER_DB_PASSWORD=12345678
```
