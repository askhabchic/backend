package dao

//func (dao *ClientDao) Get(id string) (*client.Client, error) {
//	q := `SELECT id, client_name, client_surname, birthday, gender, registration_date, address_id FROM public.client WHERE id = $1`
//
//	//r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
//
//	var cl client.Client
//	queryRow := r.psgr.QueryRow(ctx, q, id)
//	err := queryRow.Scan(&cl.ID, &cl.Name)
//	if err != nil {
//		return client.Client{}, err
//	}
//	return cl, nil
//}
