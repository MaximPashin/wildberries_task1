package db

const insertItem = `
INSERT INTO items
(chrt_id, track_number, price, rid, "name", sale, "size", total_price, nm_id, brand, status)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT (chrt_id) DO NOTHING
RETURNING chrt_id
`

const insertOrder = `
INSERT INTO orders
(order_id, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard, delivery, payment)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
ON CONFLICT (order_id) DO NOTHING
RETURNING order_id
`

const insertPayment = `
INSERT INTO payments
("transaction", request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT ("transaction") DO NOTHING
RETURNING "transaction"
`

const insertPerson = `
INSERT INTO persons
("name", phone, zip, city, address, region, email)
VALUES($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT ("name", phone) DO NOTHING
RETURNING person_id
`

const insertItemOrderRel = `
INSERT INTO ordersbaskets
("order", chrt_id)
VALUES($1, $2)
`

const getAllOrders = `
select
(
	o.order_id,
	o.track_number,
	o.entry,
	(
		per."name",
		per.phone,
		per.zip,
		per.city,
		per.address,
		per.region,
		per.email
	),
	(
		pay."transaction",
		pay.request_id,
		pay.currency,
		pay.provider,
		pay.amount,
		pay.payment_dt,
		pay.bank,
		pay.delivery_cost,
		pay.goods_total,
		pay.custom_fee
	),
	items.arr,
	o.locale,
	o.internal_signature,
	o.customer_id,
	o.delivery_service,
	o.shardkey,
	o.sm_id,
	o.date_created,
	o.oof_shard
)
from orders o
	join persons per ON o.delivery = per.person_id
	join payments pay on o.payment = pay."transaction"
	join (
		select ob."order" as id , array_agg((i.chrt_id, i.track_number, i.price, i.rid, i."name", i.sale, i."size", i.total_price, i.nm_id, i.brand, i.status)) as arr
		from ordersbaskets ob
			join items i on ob.chrt_id = i.chrt_id
		group by id
	) as items on items.id = o.order_id
`
