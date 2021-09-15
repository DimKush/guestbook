import React, { useMemo, useRef , useEffect} from 'react'
import './style.scss'
import { useTable, useFilters, usePagination, useRowSelect } from 'react-table'
import { COLUMNS } from './columns'
import { AiOutlineSearch, AiOutlineForm, AiOutlinePlusSquare, AiOutlineMinusSquare} from 'react-icons/ai'
import { BsBoxArrowInRight } from "react-icons/bs";
import { Link, Redirect} from "react-router-dom";
import "./filters-styles.scss"
import { cookies } from '../../App'
import ModalLoading from '../modal/modal-loading'
import IndeterminateCheckbox from './selector-checkbox.jsx'
import DeleteList from '../home-list-delete/delete-list'
import Modal from "../modal/modal.jsx";

const refershTable = async(listFilters) => {
	// from GO : json: invalid use of ,string struct tag, trying to unmarshal "" into int
	console.log("JSON.stringify(listFilters)",JSON.stringify(listFilters))

	listFilters.id = Number(listFilters.id)
	const responce = await fetch("http://localhost:8007/api/lists/params", {
				method: 'POST',
				credentials: 'include',
				headers : {
						"Content-type" : "application/json", 
						"Authorization" :`Bearer ${cookies.get("jwt")}`
				},
				body :JSON.stringify(listFilters)
			})
		
			const content = await responce.json();
			
			if(content.Status === "OK" && content.Result !== null ){
				return content.Result;
				//setDataTable(content.Result);
			} else if( content.Status === "Error") {
				console.log("Message");
				return null;
			}
}


export default function ListsTable({setHeaderDescript}){
	setHeaderDescript("Lists");
	const[sidebar, setSidebar] = React.useState(false);
	const[clearInput, setClearInput] = React.useState(false);
	const[dataTable, setDataTable] = React.useState([]);
	const[loadingDonut , setLoadingDonut] = React.useState(false);
	const[mpValues , setMapValues] = React.useState(new Map());
	const[timelineLoaded, setTimelineloaded] = React.useState(false);
	const[currentUser, setCurrentUser] = React.useState("");
	const[selectedRow, setSelectedRow] = React.useState({});
	const[rowIndex,setRowIndex] = React.useState(0);
	const[editItemsButtonOn, setEditItemsButtonOn] = React.useState(true);
	const[modalMsgHead, setModalMsgHead] = React.useState("");
	const[modalMsg, setModalMsg] = React.useState("");
	const[modalActive, setModalActive] = React.useState(false);

	const columns = useMemo(() => COLUMNS , []);
	
	useEffect(() => {
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

		if (!timelineLoaded) {
			(async () => {
				setLoadingDonut(true);
				const empty = {} 
				let tableData = await refershTable(empty);
				if (tableData != null) {
					setDataTable(tableData);
				}
				setLoadingDonut(false);
			}
			)();

			setTimelineloaded(true);
		}
	  }, []);

	console.log("dataTable", dataTable);
	const {
		getTableProps,
		getTableBodyProps,
		headerGroups,
		page,
		nextPage,
		previousPage,
		canNextPage,
		canPreviousPage,
		pageOptions,
		subRows,
		prepareRow,
		gotoPage,
		pageCount,
		setPageSize,
		state,
		selectedFlatRows,
	} =  useTable(
		{
			columns, 
			data: dataTable,
			initialState : {pageIndex : 0}
		},
		useFilters,
		usePagination,
		useRowSelect,
	);
	

	const arrLength = columns.length;
	const inputRef = React.useRef([]);
	const { pageIndex, pageSize } = state;
	inputRef.current = [];

	const addToRefs = (el) => { 
		if(el && !inputRef.current.includes(el)){
			inputRef.current.push(el);
		}
	}

	const[columnsFiltered, setColumnsFiltered] = React.useState(new Map());
	const updateMap = (k,v) => {
	 	setColumnsFiltered(columnsFiltered.set(k,v));
	}

	const handleClickFind = () => {
		let values = new Map();
		inputRef.current.map( elem => values.set(elem.id, elem.value) );
		setMapValues(values);

		(async () => {
			setLoadingDonut(true);
			const tableData = await refershTable(Object.fromEntries(values));
			if (tableData != null) {
				setDataTable(tableData);
			}
			setLoadingDonut(false);
		})();

		
		

		// headerGroups.map(headerGroup => { headerGroup.headers.map(column =>{
		// 	column.setFilter(mpValues.get(column.id));
		// })});
	}

	function ColumnFilter ( {column} ) {
		const handleKeyDown = (input) => {
			if (input.key === "Enter"){
				handleClickFind();
			}
		}

		return(
			<div className="form-group">
			<span>{column.id}</span>
				
				<input class="form-field"
					defaultValue={mpValues.get(column.id)}
					type="text"
					id={column.id}
					ref={addToRefs}
					onKeyDown={handleKeyDown}
					/>
			</div>
		);
	}
	
	const showSidebar = () => setSidebar(!sidebar);
	const handleClickRefresh = () => {
		inputRef.current.map(elem => elem.value = '');
	}

	const Sidebar = () => {
		return(
			<div className={ sidebar ? "SidebarFilter active" : "SidebarFilter" }>
				{!sidebar && <button className="searchClick" onClick={showSidebar}><AiOutlineSearch/></button>}
				<div className="filters-container">
				{sidebar && 
					<div className="searchButton" onClick={showSidebar}>
					<button className="searchClick active "><BsBoxArrowInRight/> </button>
						<div className="searchText">Search in table </div>
					</div>
				}
				{
					headerGroups.map(headerGroup => (
						<div {...headerGroup.getHeaderGroupProps()}>
							{
								headerGroup.headers.map(column => (
									<div {...column.getHeaderProps()} className="search-field">
										{column.canFilter ? <ColumnFilter column={column}/> : null }
									</div>
								))
							}
							
						</div>
					))
				}
				<div className="buttons-place">
					<button className="sidebar-but" onClick={handleClickRefresh}>Refresh</button>
					<button className="sidebar-but" onClick={handleClickFind}>Find</button>
				</div>

				</div>
			</div>

		);
	} 

	const handleDeleteClick = () => {
		let selectedRowList = selectedRow;
		
		(async() => {
			const delRes = await DeleteList(selectedRowList, currentUser);
			
			setModalMsgHead("Error");
			setModalMsg(delRes.Message);
			setModalActive(true);
			setTimelineloaded(false);

			setLoadingDonut(true);
			const empty = {} 
			let tableData = await refershTable(empty);
			if (tableData != null) {
				setDataTable(tableData);
			}
			setLoadingDonut(false);
		}
		)();
	}
	
	const dropMarked = () => {
		page.map(row => { row.isSelected = false; });
	}

	const enableOrDisableItemsBut = (id) => {
		console.log("enableOrDisableItemsBut" );
		(async() => {
			const responce = await fetch(`http://localhost:8007/api/lists/${id}/items/availability`, {
			  headers : { "Content-type" : "application/json",
						"Authorization" :`Bearer ${cookies.get("jwt")}`},
			  credentials : "include",
			});

			const content = await responce.json();

			if(content.Status === "OK"){
				if(content.Count > 0) {
					console.log("false");
					setEditItemsButtonOn(false);
				} else {
					console.log("false");
					setEditItemsButtonOn(true);
				}
			} else {
				// TODO: Modal error
			}
		})();
	}
	return(
	<div className="form-container">
		<Sidebar/>
	<div className={sidebar ? "form-events active" : "form-events"}>
		<div className="ControlContainer">
			<div className="butControl">
				<Link to="/lists/create">
					<button className="but-tab-hight"><AiOutlinePlusSquare/><div className="but-tab-hight-text">New List</div></button>
				</Link>
				<Link to={selectedRow.id ? `/lists/${selectedRow.id}` : `/lists` }>
					<button className="but-tab-hight"><AiOutlineForm/><div className="but-tab-hight-text">Edit List</div></button>
				</Link>
				<button className="but-tab-hight" onClick={handleDeleteClick}><AiOutlineMinusSquare/>
					<div className="but-tab-hight-text">Delete List</div>
				</button>
				<Link to={!editItemsButtonOn ? `/lists/${selectedRow.id}/` : `/lists` }>
					<button className="but-tab-hight edit" disabled={editItemsButtonOn}><AiOutlineForm/><div className="but-tab-hight-text">Edit Items</div></button>
				</Link>
			</div>
		</div>
		<table {...getTableProps()} > 
			<thead>
				{
					headerGroups.map(headerGroup => (
						<tr {...headerGroup.getHeaderGroupProps()}>
							{
								headerGroup.headers.map(column => (
									<th {...column.getHeaderProps()}>{column.render('Header')}</th>
								))
							}
							
						</tr>
					))
				}
			</thead>
			<tbody {... getTableBodyProps()}>
				{
					page.map(row => {
						prepareRow(row)
						return(
							<tr className={row.isSelected ? "trMarked" : "trNoMarked"} {...row.getRowProps({
								// on a row click action
								onClick : () => {
									page.map(rowSelected => {
										if(rowSelected !== row){
											rowSelected.isSelected = false;
										} else{
											row.isSelected = true;
										}
									});
									
									setSelectedRow(row.original);
									enableOrDisableItemsBut(row.original.id);
								},
							})}>
								{
									row.cells.map((cell) => {
										return <td {...cell.getCellProps()}>{cell.render('Cell')}</td>
									})
								}
							</tr>
						)
					})
				}
			</tbody>
		</table>
		{/* <pre>
			<code>
				{JSON.stringify(selectedRow)}
			</code>
		</pre> */}
		<div className="tableNavigator">
		<select className="form-field pagesSize" value={pageSize} onChange={e => setPageSize(Number(e.target.value))}>
				{
					[5,10,25,50].map(pageSize => (
						<option key ={pageSize} value={pageSize}>
							Show {pageSize}
						</option>
					))
				}
			</select>
			<div className="pageShow">
				<span>
					Page {' '}
					<strong>
						{ pageIndex + 1} of {pageOptions.length}
					</strong>
				</span>
				<span>
					{' '}Go to page: {' '}
					<input class="form-field pages" type="number" defaultValue={pageIndex + 1} 
					onChange={e => { 
						const pageNumber = e.target.value ? Number(e.target.value) -1 : 0;
						if (Number(e.target.value) < 1) {
							e.target.value = 1;
						}
						gotoPage(pageNumber);
					}}/>
				</span>
				<div className="tableButtons">
					<button className="but-tab-nav" onClick={() => {gotoPage(0); dropMarked(); }} disabled={!canPreviousPage}>{'<<'}</button>
					<button className="but-tab-nav" onClick={() => {previousPage(); dropMarked(); }} disabled={!canPreviousPage} >Prev</button>
					<button className="but-tab-nav" onClick={() => {nextPage(); dropMarked(); }} disabled={!canNextPage}>Next</button>
					<button className="but-tab-nav" onClick={() => {gotoPage(pageCount - 1);  dropMarked(); }} disabled={!canNextPage}>{'>>'}</button>
				</div>
			</div>
		</div>
	</div>
	<Modal active={modalActive} setActive={setModalActive} head={modalMsgHead} msg={modalMsg} isError={false}/>
	<ModalLoading active={loadingDonut}/>
	</div>

	);
}