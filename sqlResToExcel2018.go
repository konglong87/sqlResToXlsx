//把数据库查询结果另存为xlsx形式
/**
*   	@author  kl-007
*   	@param	 rows		数据库查询结果集
*	@param	 groupName	自己需求是否有用
*   	@return  string		根据自己需求	
*   	@return  err		方法执行中遇到的异常错误
 */
func SqlResToXlsx(rows *sql.Rows,groupName string)(string,error ){
	file:= xlsx.NewFile()
	sheet,_ := file.AddSheet("sheet1")//可以增加多个sheet
	row := sheet.AddRow()
	row.SetHeightCM(1)
	columns,_ :=  rows.Columns()
	for _,colName  := range columns {
		cel := row.AddCell()
		cel.Value = colName
	}
	scanArgs := make([]interface{},len(columns))
	scanRes  := make([]interface{},len(columns))
	for k,_ := range scanRes {
		scanArgs[k] = &scanRes[k]
	}
	for rows.Next() {
			row1 := sheet.AddRow()
			row1.SetHeightCM(1)
			err := rows.Scan(scanArgs...)
			if err != nil {
				fmt.Println(err)
				return "", err 
			}
			for _,v := range scanRes {
				cel := row1.AddCell()
				if v != nil {
					cel.Value = string(v.([]byte))
				}else{
					cel.Value = ""
				}
				
			}
	}
	path := ReturnPathForClose()
	err := os.MkdirAll(path, 0777)
	dt := time.Now().Format("20060102 150405")
	docName := groupName+"_"+dt[0:6]+"_计算明细.xlsx"//文件名称按自己需求
	err = file.Save(path+docName)
	if err != nil {
				fmt.Println(err)
				return "",err 
			}
	return docName,nil
 }
//