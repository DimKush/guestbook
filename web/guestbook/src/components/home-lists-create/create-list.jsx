import React, { useEffect, useState } from 'react'
import "./create-list-style.scss"
import { cookies } from "../../App";
import Modal from "../modal/modal.jsx";
import { Redirect } from 'react-router-dom';

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
	const[idCheckboxBlocked, setIdCheckboxBlocked]= useState(true);
	const[ownerCheckboxBlocked, setOwnerCheckboxBlocked] = useState(false);
	const[Owners, setOwners] = useState([]);
	const[currentUser, setCurrentUser] = useState("");
	
	const[modalMsgHead, setModalMsgHead] = React.useState("");
	const[modalMsg, setModalMsg] = React.useState("");
	const[modalActive, setModalActive] = React.useState(false);
	const[listOwner, setListOwner] = useState("");

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
			if(content.Status === "OK"){
				setOwners(content.Result);
			} else {
				// TODO : ERROR!
			}
		  }
		)();
		
		(async () => {
			const responce = await fetch("http://localhost:8007/auth/user", {
			  headers : { "Content-type" : "application/json",
						"Authorization" :`Bearer ${cookies.get("jwt")}`},
			  credentials : "include",
			});

			const content = await responce.json();

			if(content.Status === "OK"){
				setCurrentUser(content.username);
			} else {
				// TODO: Modal error
			}
		})();
		
	}, []);

	const handleCreateClick = () => {
		if(listOwner === "" && !ownerCheckboxBlocked){ 
			setModalMsg("Owner cannot be empty.");
			setModalMsgHead("Error");
			setModalActive(true);
		}

		if(titleInput.current.value === ""){ 
			console.log("log");
			setModalMsg("The title field cannot be empty.");
			setModalMsgHead("Error");
			setModalActive(true);
		}

		const CreateObjList = {
			"id" : Number(idInput.current.value),
			"owner" : ownerCheckboxBlocked ? currentUser : listOwner,
			"title" : titleInput.current.value,
			"description" : descriptionInput.current.value,
		};
		
		console.log(JSON.stringify(CreateObjList));

		(
			async() => {
				const responce = await fetch("http://localhost:8007/api/lists/create", {
				method : "POST",
				headers : { "Content-type" : "application/json",
							"Authorization" :`Bearer ${cookies.get("jwt")}`},
				credentials : "include",
				body : JSON.stringify(CreateObjList),
				});
				const content = await responce.json();
				if(content.Status === "OK"){
					setModalMsg("Record was created.");
					setModalMsgHead("OK");
					setModalActive(true);
				} else {
					setModalMsg(content.Message);
					setModalMsgHead(content.Status);
					setModalActive(true);
				}
			}
		)();
	}

	const handleCleanFieldsClick = () => {
		idInput.current.value = "";
		titleInput.current.value = "";
		descriptionInput.current.value = "";

		setOwnerCheckboxBlocked(false);
		auto_owner_checkbox.current.checked = false;

	}

	return(
		<div className="list-card-main">
			<div className="row-form">
				<h1>Create new list</h1>
			</div>
			<div className="row-form">
				<div className="search-field id checkbox">
					<input type="checkbox" id="autoId" defaultChecked={idCheckboxBlocked} ref={auto_id_checkbox} onChange={() =>
						{
							setIdCheckboxBlocked(!idCheckboxBlocked)
							idInput.current.value = "";
						}}>
						</input>
					<label for ="autoId">Auto-increment Id</label>
				</div>

				<div className="search-field id checkbox">
					<input type="checkbox" id="autoOwner" defaultChecked={ownerCheckboxBlocked} ref={auto_owner_checkbox} onChange ={() =>{
						setOwnerCheckboxBlocked(!ownerCheckboxBlocked);
						
					}}/>

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
						<select className="form-field ownerSelect" ref={ownerInput} disabled={ownerCheckboxBlocked} onChange={e => setListOwner(e.target.value)}> 
							<option disabled selected value>{ownerCheckboxBlocked ? currentUser : "-- Select an owner --"}</option>
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
				<button className="control-but" onClick={handleCleanFieldsClick}>Clean fields</button>
			</div>
			<Modal active={modalActive} setActive={setModalActive} head={modalMsgHead} msg={modalMsg} isError={false}/>
		</div>
	);
} 