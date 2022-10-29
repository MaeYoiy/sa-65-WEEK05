import React, { useEffect } from "react";

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
//
import Snackbar from "@mui/material/Snackbar";
import Divider from "@mui/material/Divider";
import FormControl from "@mui/material/FormControl";  //import มาหมดเเละเก็บไว้ในตัวแปร FormControl
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";  //import มาเฉพราะฟังก์ชัน LocalizationProvider

//import { UsersInterface } from "../models/IUser";

//import { DataGrid, GridColDef } from "@mui/x-data-grid";

//สี
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { green } from "@mui/material/colors";

//timedate
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";

import { ResearchRoomReservationRecordInterface } from "../interfaces/IResearchRoomReservationRecord";
import { ResearchRoomsInterface } from "../interfaces/IResearchRoom";

//combobox
import { ComboBoxComponent } from "@syncfusion/ej2-react-dropdowns";
import { RoomTypesInterface} from "../interfaces/IRoomType";
import { Schedule } from "@mui/icons-material";
import { AddOnsInterface} from "../interfaces/IAddOn";
import { TimeRoomsInterface} from "../interfaces/ITimeRoom";

import Select, { SelectChangeEvent } from "@mui/material/Select";
import MenuItem from '@mui/material/MenuItem';
import { UsersInterface } from "../interfaces/IUser";

const theme = createTheme({
  palette: {
    primary: {
      main: green[500],
    },
    secondary: {
      main: "#e8f5e9",
    },
  },
});

//เด้งบันทึกสำเร็จ ไม่สำเร็จ
const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,

  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});




function RecordCreate() {

  //  const [date, setDate] = React.useState<Date | null>(null);
  // useState มีไว้รับข้อมูลจากการ input เป็นหลัก
  const [researchroomreservationrecord, setResearchroomreservationrecord] = React.useState<Partial<ResearchRoomReservationRecordInterface>>({
    ResearchRoomID: 0,
    UserID: 0,
    AddOnID: 0,
    TimeRoomID: 0,
    BookDate: new Date(),
  });

  //Partial คือเลือกค่า set ค่าได้เฉพาะตัวได้
  const [success, setSuccess] = React.useState(false);

  const [error, setError] = React.useState(false);

  //เราส่งมาในรูปแบบอาเรย์ ทำการดึงข้อมูล

  //Research_Room_Reservation_Record
  const [researchroom, setResearchroom] = React.useState<ResearchRoomsInterface[]>( [] );
  const [user, setUser] = React.useState<UsersInterface[]>( [] );
  const [addon, setAddon] = React.useState<AddOnsInterface[]>( [] );
  const [time, setTime] = React.useState<TimeRoomsInterface[]>( [] );
  const [selectedDate, setSelectedDate] = React.useState<Date | null>(
    new Date()
  );

  const handleDateChange = (date: Date | null) => {
    setSelectedDate(date);
  };

  //Research_Room_Reservation_Record
  //ResearchRoom
  const getResearchroom = async () => {
    const apiUrl = `http://localhost:8080/researchrooms`;

    const requestOptions = {
      method: "GET",

      headers: {
        "Content-Type": "application/json",
      },
    };
    //การกระทำ //json
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res.data); //show ข้อมูล

        if (res.data) {
          setResearchroom(res.data);
        } else {
          console.log("else");
        }
      });
  };

  //User
  const getUser = async () => {
    const apiUrl = `http://localhost:8080/users`;

    const requestOptions = {
      method: "GET",

      headers: {
        "Content-Type": "application/json",
      },
    };
    //การกระทำ //json
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res.data); //show ข้อมูล

        if (res.data) {
          setUser(res.data);
        } else {
          console.log("else");
        }
      });
  };

  //Addon
  const getAddon = async () => {
    const apiUrl = `http://localhost:8080/addons`;

    const requestOptions = {
      method: "GET",

      headers: {
        "Content-Type": "application/json",
      },
    };
    //การกระทำ //json
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res.data); //show ข้อมูล

        if (res.data) {
          setAddon(res.data);
        } else {
          console.log("else");
        }
      });
  };

  //Time
  const getTime = async () => {
    const apiUrl = `http://localhost:8080/times`;

    const requestOptions = {
      method: "GET",

      headers: {
        "Content-Type": "application/json",
      },
    };
    //การกระทำ //json
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res.data); //show ข้อมูล

        if (res.data) {
          setTime(res.data);
        } else {
          console.log("else");
        }
      });
  };

  //เปิดปิดตัว Alert
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

  console.log(researchroomreservationrecord);

  //ทุกครั้งที่พิมพ์จะทำงานเป็น state เหมาะสำหรับกับคีย์ textfield
  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof RecordCreate;

    const { value } = event.target;

    //  setUser({ ...user, [id]: value });
  };

  //กดเลือกคอมโบไม่ได้
  const handleChange = (
    event: SelectChangeEvent<number>
  ) => {
    const name = event.target.name as keyof typeof researchroomreservationrecord;
    setResearchroomreservationrecord({
      ...researchroomreservationrecord,
      [name]: event.target.value,
    });

  };

  //
  function submit() {
    let data = {
      //กับ คอน บรรทัด 48-50 แค่ข้างหน้า ชื่อต้องตรง!!!!!!!
      ResearchRoomID: researchroomreservationrecord.ResearchRoomID,

      UserID: researchroomreservationrecord.UserID,

      AddOnID: researchroomreservationrecord.AddOnID,

      TimeID: researchroomreservationrecord.TimeRoomID,

      BookDate: researchroomreservationrecord.BookDate,

      //Date: bookdate,
    };

    console.log(data)
    const apiUrl = "http://localhost:8080/saveresearchroomreservationrecords";

    const requestOptions = {
      method: "POST",

      headers: { "Content-Type": "application/json" },

      //แปลงข้อมูลเป็นรูปแบบ {"..."}
      body: JSON.stringify(data),
      
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        if (res.data) {
          setSuccess(true);
        } else {
          setError(true);
        }
      });
  }

  //useEffect ใช้ดึงข้อมูลก่อนทำ 1 ขึ้นตอน
  useEffect(() => {
    //เอามาใช้ๆๆๆๆๆๆๆ*********
    getResearchroom();
    getUser();
    getAddon();
    getTime();
  }, []);


  return(
    <ThemeProvider theme={theme}>
      <Container maxWidth="md">
        <Snackbar
          open={success}
          autoHideDuration={6000}
          onClose={handleClose}
          anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
        >
          <Alert onClose={handleClose} severity="success">
            บันทึกข้อมูลสำเร็จ
          </Alert>
        </Snackbar>

        <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="error">
            บันทึกข้อมูลไม่สำเร็จ
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
                color="inherit"
                gutterBottom
              >
                จองห้องค้นคว้า
              </Typography>
            </Box>
          </Box>

          <Divider />

          <Grid container spacing={3} sx={{ padding: 2 }}>

            {/*room type 
            <Grid item xs={6}>
              <p>Room Type</p>

              <FormControl fullWidth variant="outlined">
                <Select
                  value = {researchroomreservationrecord.RoomTypeID}
                  onChange = {handleChange}
                  inputProps={{
                    nameResearch: "RoomTypeID",
                  }}
                    // defaultValue={0}
                >
                  <MenuItem value={0} key={0}>
                    เลือกประเภทห้องค้นคว้า
                  </MenuItem>
                  {roomtype.map((item: RoomTypesInterface) => (
                    <MenuItem value={item.ID}>{item.Type}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            {/*Equipment 
            <Grid item xs={6}>
              <p>Equipment</p>

              <FormControl fullWidth variant="outlined">
                <Select
                  value = {researchroomreservationrecord.EquipmentID}
                  onChange = {handleChange}
                  inputProps={{
                    nameResearch: "EquipmentID",
                  }}
                    // defaultValue={0}
                >
                  <MenuItem value={0} key={0}>
                    เลือกอุปกรณ์สำหรับห้อง
                  </MenuItem>
                  {equipment.map((item: EquipmentsInterface) => (
                    <MenuItem value={item.ID}>{item.Name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
                  </Grid> */}

            {/*Addon*/}
            <Grid item xs={6}>
              <p>Add-On</p>

              <FormControl fullWidth variant="outlined">
                <Select
                  value = {researchroomreservationrecord.AddOnID}
                  onChange = {handleChange}
                  inputProps={{
                    name: "AddOnID",
                  }}
                  // defaultValue={0}
                >
                  <MenuItem value={0} key={0}>
                    เลือกอุปกรณ์เสริม
                  </MenuItem>
                  {addon.map((item: AddOnsInterface) => (
                    <MenuItem value={item.ID}>{item.Name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            {/*<Grid item xs={6}></Grid>*/}

            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>Time</p>

                <Select
                  value={researchroomreservationrecord.TimeRoomID}
                  onChange={handleChange}
                  inputProps={{
                    name: "TimeID",
                  }}
                  // defaultValue={0}
                >
                  <MenuItem value={0} key={0}>
                    เลือกเวลา
                  </MenuItem>
                  {time.map((item: TimeRoomsInterface) => (
                    <MenuItem value={item.ID}>{item.Period}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            {/* //วันที่และเวลา 
            <Grid item xs={7}>
              <FormControl fullWidth variant="outlined">
                <p>วันที่และเวลา</p>

                <LocalizationProvider dateAdapter={AdapterDayjs}>
                  <DateTimePicker
                    renderInput={(props) => <TextField {...props} />}
                    label="กรุณาเลือกวันและเวลา"
                    value={selectedDate} //แก้
                    // onChange={(newValue) => {
                    // setDate(newValue);

                    // }}
                    onChange={setSelectedDate}
                  />
                </LocalizationProvider>
              </FormControl>
                  </Grid>*/}

            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>วันที่และเวลา</p>
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
              <Button component={RouterLink} to="/" variant="contained">
                <Typography
                  color="secondary"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
                  ย้อนกลับ
                </Typography>
              </Button>

              <Button
                style={{ float: "right" }}
                onClick={submit}
                variant="contained"
                color="primary"
              >
                <Typography
                  color="secondary"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
                  จองห้องค้นคว้า
                </Typography>
              </Button>
            </Grid>
          </Grid>
        </Paper>
      </Container>
    </ThemeProvider>
  )
    
        {/*<div>
          <Container maxWidth="md">
                <Box sx={{ bgcolor: '#cfe8fc', height: '100vh' }} />
                <Paper>
                    <Box display={"flex"} sx={{marginTop: 2, paddingX: 2, paddingY: 2, bgcolor: '#cfe8fc' }}>
                        <Button>Save</Button>
                        <h2>SHOW RECORD</h2>
                    </Box>
                    <hr />
                    <Grid container spacing={2}>
                        <Grid item xs={7}>
                            <Autocomplete
                                disablePortal
                                id="combo-box-demo"
                                options={top100Films}
                                sx={{ width: 300 }}
                                renderInput={(params) => <TextField {...params} label="Roomtype" />}
                            />
                            
                        </Grid>
                        <Grid item xs={4}>
                            
                        </Grid>
                        <Grid item xs={4}>
                            
                        </Grid>
                        <Grid item xs={8}>
                            
                        </Grid>
                    </Grid>                    
                </Paper>
  
                
                
              </Container>

        </div>*/}
  


}

// Top 100 films as rated by IMDb users. http://www.imdb.com/chart/top
const top100Films = [
    { label: 'The Shawshank Redemption', year: 1994 },
    { label: 'The Godfather', year: 1972 },
    { label: 'The Godfather: Part II', year: 1974 },
  ];
     
        


export default RecordCreate;

