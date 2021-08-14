import React from "react";
import "./filters-styles.scss"

export const ColumnFilter = ({ column }) => {
	const {filterValue, setFilter} = column
	console.log(column);
	return (
		<div className="form-group">
		<span>{column.id}</span>
			
			<input class="form-field"	
				//value={filterValue || ''} 
				//onChange ={(event) => setFilter(event.target.value)} />
				/>
		</div>
	);
}