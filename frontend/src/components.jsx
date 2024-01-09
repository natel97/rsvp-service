import { styled } from "styled-components";
import moment from "moment";

export const PageHeader = styled.h1`
  background: #088;
  padding: 2rem;
  text-align: center;
  font-size: 40px;
`;

export const Card = styled.div`
  padding: 16px;
  margin: 32px 16px;
  text-align: left;
  border-radius: 8px;
  box-shadow: 0px 0px 4px 4px rgba(0, 0, 0, 0.25);
  background: #046;
`;

export const Action = styled.button`
  text-align: left;
  padding: 12px;
  font-size: 1.5rem;
  margin: 4px 16px;
  background: #161616;
`;

export const Option = ({
  prefix,
  options = [],
  current = "",
  setCurrent = () => {},
}) => {
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "space-between",
        margin: "16px",
      }}
    >
      {options.map((option) => (
        <span key={option}>
          <input
            id={`${prefix}-${option}`}
            name={prefix}
            type="radio"
            value={option}
            checked={option === current}
            onChange={() => setCurrent(option)}
            style={{ display: "none" }}
          />
          <label htmlFor={`${prefix}-${option}`}>{option}</label>
        </span>
      ))}
    </div>
  );
};

const Attendance = ({ yes, no, maybe }) => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "row",
        justifyContent: "space-between",
        marginTop: "48px",
      }}
    >
      <div>🟢 Yes ({yes})</div>
      <div>🔴 No ({no})</div>
      <div>🤷‍♂️ Maybe ({maybe})</div>
    </div>
  );
};

const EmojiDetail = ({ emoji = "", children, size = "1rem" }) => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "row",
        alignItems: "center",
        margin: "16px 0",
      }}
    >
      <div style={{ fontSize: size, marginRight: "8px" }}>{emoji}</div>
      <div>{children}</div>
    </div>
  );
};

export const EventCard = ({
  title = "Game Night",
  date = Date.parse("23 Jan 2023 18:30:00 GMT+1100"),
  street = "111 Flinders Street",
  city = "Melbourne, VIC 3000",
  attendance = { yes: 3, no: 2, maybe: 8 },
}) => {
  const day = moment(date).format("MMMM Do YYYY");
  const time = moment(date).format("hh:mm A");

  return (
    <Card>
      <h2 style={{ fontSize: "2rem", textAlign: "center" }}>{title}</h2>

      <EmojiDetail size="2rem" emoji="📅">
        <h3>{day}</h3>
        <h3>{time}</h3>
      </EmojiDetail>

      <EmojiDetail size="2rem" emoji="📍">
        <h3>{street}</h3>
        <h3>{city}</h3>
      </EmojiDetail>
      {attendance && <Attendance {...attendance} />}
    </Card>
  );
};
