import { useEffect, useState } from 'react';


function Notification({ message, onClose }) {
  const [show, setShow] = useState(true);

  useEffect(() => {
    if (message) {
      const timer = setTimeout(() => {
        setShow(false);
        onClose(); // Call the onClose function to remove the notification
      }, 3000); // Display for 3 seconds

      return () => clearTimeout(timer); // Cleanup timer on component unmount
    }
  }, [message, onClose]);

  if (!show) return null;

  return (
    <div className="notification">
      {message}
    </div>
  );
}

export default Notification;
