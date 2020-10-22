import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [drugs, setDrugs] = useState(Array);
 const [loading, setLoading] = useState(true);

 useEffect(() => {
   const getDrugs = async () => {
     const res = await api.listDrug({ limit: 10, offset: 0 });
     setLoading(false);
     setDrugs(res);
   };
   getDrugs();
 }, [loading]);
 
 const listEmployees = drugs.map((name) =>
 <li>{drugs.map.name}</li>
);

 const deleteDrugs = async (id: number) => {
   const res = await api.deleteDrug({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">Employee</TableCell>
           <TableCell align="center">Name</TableCell>
           <TableCell align="center">Drugtype</TableCell>
           <TableCell align="center">Property.</TableCell>
           <TableCell align="center">Howto</TableCell>
           <TableCell align="center">Disease</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {drugs.map((item:any) => (
           <TableRow key={item.id}>
             
             <TableCell align="center">{item.edges?.employee?.name}</TableCell>
             <TableCell align="center">{item.name}</TableCell>
             <TableCell align="center">{item.edges?.drugtype?.name}</TableCell>
             <TableCell align="center">{item.property}</TableCell>
             <TableCell align="center">{item.howto}</TableCell>
             <TableCell align="center">{item.edges?.disease?.name}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteDrugs(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
