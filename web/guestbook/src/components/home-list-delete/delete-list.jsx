import React, { useState } from 'react'
import { cookies } from "../../App";
import Modal from "../modal/modal.jsx";


export default function DeleteList(selectedRow) {
	let currentUser;
	let result = {Status : "", Message : ""};
	
	console.log(selectedRow);
	
	let responce = new Promise((resolve, reject) => {  
		fetch("http://localhost:8007/auth/user", {
		headers : { "Content-type" : "application/json",
				"Authorization" :`Bearer ${cookies.get("jwt")}`},
		credentials : "include",
	}).then(responce => responce.json())
	  .then(json => resolve((json.username)))
	
	});

	currentUser = responce.then(username => username);


	console.log("R " , currentUser);
	if(currentUser !== selectedRow.owner){
		result.Status = "Error";
		result.Message = `Access denied. You don't have permissions to delete with list. Owner of the record is ${selectedRow.owner}.`;

		return result;
	}

	return "OK";
}