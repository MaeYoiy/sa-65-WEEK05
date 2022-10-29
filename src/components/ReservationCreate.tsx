import React, { useEffect, useState } from "react";

import { Link as RouterLink } from "react-router-dom";

import Container from "@mui/material/Container";
import MuiAlert, { AlertProps } from "@mui/material/Alert";

import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';

import Snackbar from "@mui/material/Snackbar";
import Divider from "@mui/material/Divider";
import FormControl from "@mui/material/FormControl";  //import มาหมดเเละเก็บไว้ในตัวแปร FormControl
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";  //import มาเฉพราะฟังก์ชัน LocalizationProvider

//timedate
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";

import { ResearchRoomReservationRecordInterface } from "../interfaces/IResearchRoomReservationRecord";
import { ResearchRoomsInterface } from "../interfaces/IResearchRoom";

import { RoomTypesInterface} from "../interfaces/IRoomType";
import { AddOnsInterface} from "../interfaces/IAddOn";
import { TimeRoomsInterface} from "../interfaces/ITimeRoom";

import Select, { SelectChangeEvent } from "@mui/material/Select";
import MenuItem from '@mui/material/MenuItem';
import { UsersInterface } from "../interfaces/IUser";

import {
  GetRoomTypes,
  GetResearchRooms,
  GetAddOns,
  GetTimeRooms,
  ResearchRoomReservationRecords,
} from "../services/HttpClientService";
const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function ReservationCreate() {
  // useState มีไว้รับข้อมูลจากการ input เป็นหลัก
  //เราส่งมาในรูปแบบอาเรย์ ทำการดึงข้อมูล
  //ResearchRoom
  //const [roomtypes, setRoomtypes] = useState<RoomTypesInterface[]>([]);
  //ResearchRoomReservationRecord
  const [researchrooms, setResearchrooms] = useState<ResearchRoomsInterface[]>([]);
  const [addons, setAddons] = useState<AddOnsInterface[]>([]);
  const [timerooms, setTimerooms] = useState<TimeRoomsInterface[]>([]);
  const [researchroomreservationrecord, setResearchroomreservationrecord] = useState<ResearchRoomReservationRecordInterface>({
    BookDate: new Date()
  });

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  //เปิดปิดตัว alert
  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  //combobox
  //set state ของ researchroomreservationrecord หรือ set setResearchroomreservationrecord
  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof researchroomreservationrecord;
    setResearchroomreservationrecord({
      ...researchroomreservationrecord,
      [name]: event.target.value,
    });
  };

  const getResearchrooms = async () => {
    let res = await GetResearchRooms();
    if (res) {
      setResearchrooms(res);
    }
  };

  const getAddons = async () => {
    let res = await GetAddOns();
    if (res) {
      setAddons(res);
    }
  };

  const getTimeRooms = async () => {
    let res = await GetTimeRooms();
    if (res) {
      setTimerooms(res);
    }
  };

  useEffect(() => {
    //getEquipments();
    //getRoomtypes();
    getResearchrooms();
    getAddons();
    getTimeRooms();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      ResearchRoomID: convertType(researchroomreservationrecord.ResearchRoomID),
      UserID: convertType(localStorage.getItem("uid") as string),
      AddOnID: convertType(researchroomreservationrecord.AddOnID),
      TimeRoomID: convertType(researchroomreservationrecord.TimeRoomID),
      BookDate: researchroomreservationrecord.BookDate,
    };

    let res = await ResearchRoomReservationRecords(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }

  return(
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          จองห้องค้นคว้าสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          จองไม่สำเร็จ ห้องนี้ในช่วงเวลานี้ถูกจองเเล้ว
        </Alert>
      </Snackbar>
      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              จองห้องค้นคว้า
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Add On</p>
              <Select
                native
                value={researchroomreservationrecord.AddOnID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "AddOnID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกอุปกรณ์เสริม
                </option>
                {addons.map((item: AddOnsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Time Access</p>
              <Select
                native
                value={researchroomreservationrecord.TimeRoomID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "TimeRoomID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกเวลา
                </option>
                {timerooms.map((item: TimeRoomsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Period}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          {/* ResearchRooms */}
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Research Room</p>
              <Select
                native
                value={researchroomreservationrecord.ResearchRoomID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "ResearchRoomID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกห้องค้นคว้า
                </option>
                {researchrooms.map((item: ResearchRoomsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Select Date</p>
              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={researchroomreservationrecord.BookDate}
                  onChange={(newValue) => {
                    setResearchroomreservationrecord({
                      ...researchroomreservationrecord,
                      BookDate: newValue,
                    });
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/researchroomreservationrecords"
              variant="contained"
              color="inherit"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              // component={RouterLink}
              // to="/researchroomreservationrecords"
              variant="contained"
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default ReservationCreate;