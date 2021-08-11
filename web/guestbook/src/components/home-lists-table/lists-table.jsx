import React, { useMemo } from 'react'
import './style.scss'
import { useTable } from 'react-table'
import MOCK_DATA from './MOCK_DATA.json'
import { COLUMNS } from './columns'

export default function ListsTable({setHeaderDescript}){
	setHeaderDescript("Lists");
	const columns = useMemo(() => COLUMNS , []);
	const data = useMemo(() => MOCK_DATA, []);
	
	//prepare table
	const TableInstance = useTable({
		columns,
		data
	});

	const {
		getTableProps,
		getTableBodyProps,
		headerGroups,
		rows,
		prepareRow,
	} =  TableInstance;

	return(
	<div className="form-events">
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
					rows.map(row => {
						prepareRow(row)
						return (
							<tr {...row.getRowProps()}>
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
	</div>
	);
}