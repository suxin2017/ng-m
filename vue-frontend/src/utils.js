import moment from "moment";

export function formatTime(row, column) {
  const date = row[column.property];
  if (date == undefined) {
    return "";
  }
  return formatDate(date);
}

export const formatDate = (date) => {
  return moment(date).format("YYYY-MM-DD HH:mm:ss");
};
