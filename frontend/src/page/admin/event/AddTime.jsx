import { useState } from "react";
import {
  ActionButton,
  Modal,
  PageHeader,
  TextInput,
} from "../../../components";
import { getAuthToken } from "../auth";
import { getDefaultDateTime } from "./create";
import moment from "moment";

const addTimeOption = (id, Time) => {
  const apiKey = getAuthToken();
  return fetch(
    `${import.meta.env.VITE_API_URL}/admin/event/${id}/time-option`,
    {
      method: "POST",
      body: JSON.stringify({
        Time: moment(Time).toDate(),
      }),
      headers: {
        Authorization: apiKey,
      },
    }
  );
};

const AddTime = ({ eventID = "" }) => {
  const [open, setOpen] = useState(false);
  const [date, setDate] = useState(getDefaultDateTime());

  return (
    <>
      <ActionButton onClick={() => setOpen(true)}>Add Time Option</ActionButton>
      <Modal open={open} setIsOpen={setOpen}>
        <PageHeader>Add Time</PageHeader>
        <TextInput
          type="datetime-local"
          placeholder="Date"
          value={date}
          min={new Date()}
          onChange={(e) => setDate(e.target.value)}
        />
        <ActionButton onClick={() => addTimeOption(eventID, date)}>
          Create
        </ActionButton>
      </Modal>
    </>
  );
};

export default AddTime;
