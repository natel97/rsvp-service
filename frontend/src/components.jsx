import { styled } from "styled-components";
import moment from "moment";

const colors = {
  success: {
    light: "#AF7",
    dark: "#053",
  },
  primary: {
    light: "#6CC",
    dark: "#088",
  },
  secondary: {
    light: "#BDF",
    dark: "#046",
  },
};

export const TextInput = styled.input`
  border: 1px solid #ccc;
  width: 80%;
  margin: 16px 0;
  padding: 12px;
  font-size: 1.5rem;
  border-radius: 8px;
`;

export const Radio = styled.input`
  & + label {
    box-shadow: 0px 0px 2px 2px rgba(255, 255, 255, 0.25);
    border-radius: 8px;
    padding: 8px;
  }

  &:checked + label {
    background: ${colors.success.dark};
  }

  &:checked + label {
    @media (prefers-color-scheme: light) {
      background: ${colors.success.light};
    }
  }
`;

export const PageHeader = styled.h1`
  background: ${colors.primary.dark};
  padding: 1.3rem;
  text-align: center;
  font-size: 38px;

  @media (prefers-color-scheme: light) {
    background: ${colors.primary.light};
  }
`;

export const ActionButton = styled.button`
  font-size: 1.5rem;
  background: ${colors.success.dark};
  margin: 8px 0;
  padding: 1rem;
  border-radius: 8px;

  @media (prefers-color-scheme: light) {
    background: ${colors.success.light};
  }
`;

export const Card = styled.div`
  padding: 16px;
  margin: 32px 16px;
  text-align: left;
  border-radius: 8px;
  box-shadow: 0px 0px 4px 4px rgba(0, 0, 0, 0.25);
  background: ${colors.secondary.dark};

  @media (prefers-color-scheme: light) {
    background: ${colors.secondary.light};
  }
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
          <Radio
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
      <div>🤷‍♂️ Maybe ({maybe})</div>
      <div>🔴 No ({no})</div>
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

export const DAY_FORMAT = "MMMM Do YYYY";
export const TIME_FORMAT = "hh:mm A";

export const EventCard = ({
  title = "- - -",
  date = Date.parse("23 Jan 2023 18:30:00 GMT+1100"),
  street = "- - -",
  city = "- - - -, - - -  0000",
  attendance = { yes: -1, no: -1, maybe: -1 },
  description = "-----",
  onClick,
}) => {
  const day = moment(date).format(DAY_FORMAT);
  const time = moment(date).format(TIME_FORMAT);

  return (
    <Card onClick={onClick}>
      <h2 style={{ fontSize: "2rem", textAlign: "center" }}>{title}</h2>

      <EmojiDetail size="2rem" emoji="📅">
        <h3>{day}</h3>
        <h3>{time}</h3>
      </EmojiDetail>

      <EmojiDetail size="2rem" emoji="📍">
        <h3>{street}</h3>
        <h3>{city}</h3>
      </EmojiDetail>

      <div>{description}</div>
      {attendance && <Attendance {...attendance} />}
    </Card>
  );
};
