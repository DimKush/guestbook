import React from 'react'
import "./create-list-style.scss"

export default function CreateList() { 
	console.log("newList");
	function ColumnCreateList ( {column} ) {
		return(
			<div className="form-group">
			<span>{column}</span>
				
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
			<div className="row-form">
				<h1>Create new list</h1>
			</div>
			<div className="row-form">
				<div className="search-field id">
					<ColumnCreateList column={"Id"}/>
				</div>
				<div className="search-field title" >
					<ColumnCreateList column={"Title"}/>
				</div>
				<div className="search-field owner" >
					<ColumnCreateList column={"Owner"}/>
				</div>
			</div>
			<div className="row-form">
				<div className="search-field id checkbox">
					<input type="checkbox" id="autoId"></input>
					<label for ="autoId">Auto-increment Id</label>
				</div>
				<div className="search-field id checkbox">
					<input type="checkbox" id="autoTitle"></input>
					<label for ="autoTitle">Autogenerate title</label>
				</div>
				<div className="search-field id checkbox">
					<input type="checkbox" id="autoOwner"></input>
					<label for ="autoOwner">I'm the owner</label>
				</div>
			</div>
			<div className="row-form">
				<div className="search-field description">
					<ColumnCreateList column={"Description"}/>
				</div>
			</div>
			<div className="row-form">
				<button className="control-but">Create List</button>
				<button className="control-but">Clean fields</button>
			</div>
		</div>
	);
} 