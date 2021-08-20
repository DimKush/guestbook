import React from 'react'


export default function CreateList() { 
	console.log("newList");
	function ColumnCreateList ( {column} ) {
		return(
			<div className="form-group">
			<span>{column.id}</span>
				
				<input class="form-field"
					type="text"
					id={column}
					//ref={addToRefs}
					//onKeyDown={handleKeyDown}
					/>
			</div>
		);
	}

	return(
		<div className="list-card-main">
			<div className="search-field">
				fsgfsdgdfgdfgdfgdgdfsaasdasdgdfgdgdgfda
			</div>

		</div>
	);
} 