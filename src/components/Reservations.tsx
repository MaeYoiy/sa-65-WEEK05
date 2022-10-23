import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
// import { WatchVideoInterface } from "../interfaces/IWatchVideo";
import { ResearchRoomReservationRecordInterface } from "../interfaces/IResearchRoomReservationRecord";
import { GetResearchRoomReservationRecords } from "../services/HttpClientService";
import { GetResearchRooms } from "../services/HttpClientService";

function Reservations() {
  const [researchroomreservationrecords, setResearchroomreservationrecords] = useState<ResearchRoomReservationRecordInterface[]>([]);

  useEffect(() => {
    getResearchroomreservationrecords();
  }, []);

  const getResearchroomreservationrecords = async () => {
    let res = await GetResearchRoomReservationRecords();
    if (res) {
      setResearchroomreservationrecords(res);
    } 
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 100 },
    {
      field: "User",
      headerName: "ผู้ใช้",
      width: 150,
      valueFormatter: (params) => params.value.FirstName,
    },
    {
      field: "AddOn",
      headerName: "อุปกรณ์เสริม",
      width: 200,
      valueFormatter: (params) => params.value.Name,
    },
    {
      field: "ResearchRoom",
      headerName: "ประเภทห้อง",
      width: 150,
      valueFormatter: (params) => params.value.RoomType.Type,
    },
    {
      field: "TimeRoom",
      headerName: "เวลาจอง",
      width: 150,
      valueFormatter: (params) => params.value.Period,
    },
    { field: "BookDate", headerName: "วันที่", width: 110 },
    
  ];

  return (
    <div>
      <Container maxWidth="md">
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลการจองห้องค้นคว้า
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/researchroomreservationrecords/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={researchroomreservationrecords}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={6}
            rowsPerPageOptions={[6]}
          />
        </div>
      </Container>
    </div>
  );
}

export default Reservations;