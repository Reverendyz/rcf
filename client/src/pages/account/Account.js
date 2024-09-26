import api from "../../api/api";
import { useEffect, useState } from "react";
import ParticipantsTable from "../../components/table/ParticipantsTable";

function Account() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await api.get('/participants');
        setData(response.data);
      } catch (err) {
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error.message}</div>;

  const participants = data.participants.length > 0 ? data.participants : "Sem contas cadastradas"
  return (
    <div>
      <ParticipantsTable data={participants} />
    </div>
  );
}

export default Account;
