
export default function ServerStatus(setModalMsg, setIsError){
	fetch("http://localhost:8007/status", {
			method: 'GET',
			headers : {
				'Content-Type' : 'application/json'
			}
		}).then(responce => responce.json()).then(data => {
			setModalMsg(data.Message);
			setIsError(false);
		}).catch(error => {setModalMsg("Server is dead."); setIsError(true); });
}