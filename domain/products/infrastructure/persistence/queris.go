package persistence

const (
	insertProduct  = "INSERT INTO public.productos (product_id,product_name,product_amount,product_user_created,product_date_created,product_user_modify,product_date_modify) VALUES ($1,$2,$3,$4,'NOW()',$5,'NOW()')"
	GetAllProducts = "SELECT * FROM public.productos"
)
