# Trading Bot Strategy

Factors:

    Time under trend:
        --- How many time it's on the upper or under trend per (hour,minute,day).
    Average of the stock/cryptocurrency:
        --- Given a time window get the average of the stock and compare.
    Trend Size:
        --- How big or low it's the trend compared to the previous aka volatility.

Orders:

    Invest/trend correlation:
        -- Given the volatility and the trend size this decides how we place a big or low the percentage of the order from a budget.
    Stop-Loss: 
        -- Based on a target/average/trend calculate stop-loss -- sell but keep track
            * If the stock recovers till the limit buy again.

TODOS:

    1. Take historical data from 1 year analyse the timer under trend and have a constant.
    2. Update that constant accordingly ot the following prices.
    3. Take the average trend of a window(1 day) store those in a database -- 10 averages.
    4. Calculate trend size aka volatility between trends and trends.  
