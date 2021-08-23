import React, { useEffect, useState } from 'react'
import "./create-list-style.scss"
import { cookies } from "../../App";

function ColumnCreateList ( {column, ref_current, blocked=false} ) {
	return(
		<div className="form-group">
		<span>{column}</span>
			
			<input class="form-field"
				type="text"
				id={column}
				ref={ref_current}
				disabled={blocked}
				/>
		</div>
	);
}


export default function CreateList() {
	const[idCheckboxBlocked, setIdCheckboxBlocked]= useState(false);
	const[ownerCheckboxBlocked, setOwnerCheckboxBlocked] = useState(false);
	const[Owners, setOwners] = useState([])
	let idInput = React.createRef();
	let ownerInput = React.createRef();
	let titleInput = React.createRef();
	let descriptionInput = React.createRef();
	let auto_id_checkbox = React.createRef();
	let auto_owner_checkbox = React.createRef();

	useEffect(() => {
		(
		  async () => {
			const responce = await fetch("http://localhost:8007/api/users/GetAllUsernames", {
			  headers : {"Content-type" : "application/json",
						 "Authorization" :`Bearer ${cookies.get("jwt")}`},
			  credentials : "include",
			  
			});
			
			const content = await responce.json();
	
			console.log(cookies);
			if(content.Status === "OK"){
				setOwners(content.Result)
			  //setAuthStatus(true);
			} else {
			  //setAuthStatus(false);
			}
		  }
		)();
	
	});

	const handleCreateClick = () => {
		const obj = {
			"id" : idInput.current.value,
			"owner" : ownerInput.current.value,
			"title" : titleInput.current.value,
			"description" : descriptionInput.current.value,
			"auto_id_checkbox" : auto_id_checkbox.current.value,
			"auto_owner_checkbox" : auto_owner_checkbox.current.value,
		}
		
		console.log(JSON.stringify(obj));
	}

	return(
		<div className="list-card-main">
			<div className="row-form">
				<h1>Create new list</h1>
			</div>
			<div className="row-form">
				<div className="search-field id checkbox">
					<input type="checkbox" id="autoId" defaultValue="off" ref={auto_id_checkbox} onChange={() =>
						setIdCheckboxBlocked(!idCheckboxBlocked)}></input>
					<label for ="autoId">Auto-increment Id</label>
				</div>
				<div className="search-field id checkbox">
					
					<input type="checkbox" id="autoOwner"  defaultValue="off" ref={auto_owner_checkbox} onChange={()=>
					setOwnerCheckboxBlocked(!ownerCheckboxBlocked)}></input>

					<label for ="autoOwner">I'm the owner</label>
				</div>
				<div className="search-field id">
					<ColumnCreateList column={"Id"} ref_current={idInput} blocked={idCheckboxBlocked}/>
				</div>
			</div>
			<div className="row-form">
				
				<div className="search-field title" >
					<ColumnCreateList column={"Title"} ref_current={titleInput}/>
				</div>
				<div className="search-field owner" >
				<div className="form-group">
					<span>Owner</span>
						<select className="form-field pagesSize" disabled={ownerCheckboxBlocked}>
						{
							Owners.map(Owner => (
								<option value={Owner}>
									{Owner}
								</option>
							))
						}
			</select>
					</div>
				</div>
			</div>
			<div className="row-form">
				<div className="search-field description">
					<ColumnCreateList column={"Description"} ref_current={descriptionInput}/>
				</div>
			</div>
			<div className="row-form">
				<button className="control-but" onClick={handleCreateClick}>Create List</button>
				<button className="control-but">Clean fields</button>
			</div>
		</div>
	);
} 