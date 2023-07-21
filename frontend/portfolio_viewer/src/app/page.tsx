"use client"
import { Chart, registerables } from "chart.js";
import { useState, ChangeEvent } from 'react';
import styles from "./styles.module.css"

type Dictionary<value> = { [key: string]: value };

type Stock = {
  Name: string;
  Ticker: string;
  Value: number;
  Price: number;
  Shares: number;
}

type Sector = {
  Sector: string;
  Value: number;
  Stocks: Stock[];
}

type Portfolio = {
  TotalValue: number;
  Sectors: Sector[];
  SectorPercentage: Dictionary<number>;
}
Chart.register(...registerables)
const LABELS = [
  "BONDS AND EQ",
  "ENERGY",
  "MATERIALS",
  "INDUSTRIALS",
  "CONSUMER DISCRETIONARY",
  "CONSUMER STAPLES",
  "HEALTH CARE",
  "FINANCIALS",
  "INFORMATION TECHNOLOGY",
  "COMMUNICATION SERVICES",
  "UTILITIES",
  "REAL ESTATE",
];
  const COLORS: Dictionary<string> = {
    "BONDS AND EQ":"rgb(124, 112, 3)",
    "ENERGY": "rgb(63, 81, 181)",
    "MATERIALS": "rgb(33, 150, 243)",
    "INDUSTRIALS":"rgb(76, 175, 80)",
    "CONSUMER DISCRETIONARY":"rgb(255, 193, 7)",
    "CONSUMER STAPLES":"rgb(255, 87, 34)",
    "HEALTH CARE":"rgb(156, 39, 176)",
    "FINANCIALS":"rgb(0, 188, 212)",
    "INFORMATION TECHNOLOGY":"rgb(255, 152, 0)",
    "COMMUNICATION SERVICES":"rgb(103, 58, 183)",
    "UTILITIES":"rgb(0, 150, 136)",
    "REAL ESTATE":"rgb(255, 235, 59)",
  };

  var sectorData: Dictionary<number> = {
    "BONDS AND EQ": 0,
    "ENERGY": 0,
    "MATERIALS": 0,
    "INDUSTRIALS": 0,
    "CONSUMER DISCRETIONARY": 0,
    "CONSUMER STAPLES": 0,
    "HEALTH CARE": 0,
    "FINANCIALS": 0,
    "INFORMATION TECHNOLOGY": 0,
    "COMMUNICATION SERVICES": 0,
    "UTILITIES": 0,
    "REAL ESTATE": 0,
  };
var doughnutSectorData: number[] = [0,0,0,0,0,0,0,0,0,0,0,0]
var responseData: Portfolio = {
  TotalValue: 0,
  Sectors: [],
  SectorPercentage: {},
}

const Home = () => {
 const [selectedFile, setSelectedFile] = useState<File | null>(null);
 const [sectors, setSectors] = useState<Sector[]>([]);

 const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
  const file = event.target.files?.[0];
  setSelectedFile(file || null);
};

  const handleFileUpload = async () => {
    if (!selectedFile) {
      return;
    }
    try {
      const formData = new FormData();
      formData.append('file', selectedFile);
      const response = await fetch('http://localhost:8080/portfolio', {
        method: 'POST',
        body: formData,
      });

      if (response.ok) {
        // File uploaded successfully
        responseData = await response.json();
        setSectors(responseData.Sectors)
        for (const sector in responseData['SectorPercentage'] ){
          sectorData[sector] = responseData['SectorPercentage'][sector]
        }
        doughnutSectorData = UnpackSectorData(sectorData);
        const _doughnutChart = DoughnutChart(doughnutSectorData);
      } else {
        // Error uploading file
        console.error('Error uploading file:', response.statusText);
      }
    } catch (error) {
      console.error('Error uploading file:', error);
    }
  };
  return (
    <div>
      <div>
        <div className={styles.head}>
          <h1>Portfolio</h1>
          <input type="file" onChange={handleFileChange} />
          <button onClick={handleFileUpload} disabled={!selectedFile}>Submit</button>
        </div>
        {/* Doughnut chart */}
        <div className={styles.main1}>
          <div>
              <canvas id='myChart' height="600px"></canvas>
          </div>
          <div key="table" className={styles.section}>
            <TableRows Sectors={sectors} TotalValue={0} SectorPercentage={{}}/>
          </div>
        </div>
      </div>
    </div>
  )
}

function DoughnutChart(data: number[]) {
  var canvas: any = document.getElementById('myChart');
  var ctx = canvas.getContext('2d');
  if (ctx) {
    var myChart = new Chart(ctx, {
      type: 'doughnut',
      data: {
        labels: LABELS,
        datasets: [
          {
            data,
            borderColor: UnpackColors(COLORS),
            backgroundColor: UnpackColors(COLORS),
            borderWidth: 1,
          },
        ],
      },
      options: {
        maintainAspectRatio: false,
        plugins: {
          legend: {
            title: {
              display: true,
              text: "SECTORS",
              font: {
                style: "oblique",
              },
              color: 'white',
              padding: 10,
            },
            labels: {
              color: 'white',
              padding: 10,
            },
            position: "bottom",
          },
          tooltip: {
              callbacks: {
                label: function (context: any) {
                  var label = ' Percent: '
                  label += context.raw.toFixed(2) + '%';
                  const foundSector = responseData.Sectors.find((obj: Sector) => obj.Sector === context.label);
                  if (foundSector) {
                    label +='  Value: $' + foundSector.Value.toFixed(2);
                  }
                  return label;
                },
              },
          },
        },
      },
    })
    return myChart;
  }
}

function UnpackSectorData(dictionary: Dictionary<number>) {
  var list:number[] = []
  for (const [key, value] of Object.entries(dictionary)){
    list.push(value)
  }
  return list
}

function UnpackColors(dictionary: Dictionary<string>) {
  var list:string[] = []
  for (const [key, value] of Object.entries(dictionary)){
    list.push(value)
  }
  return list
}

const formatter = new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',

  // These options are needed to round to whole numbers if that's what you want.
  //minimumFractionDigits: 0, // (this suffices for whole numbers, but will print 2500.10 as $2,500.1)
  //maximumFractionDigits: 0, // (causes 2500.99 to be printed as $2,501)
});

const TableRows: React.FC<Portfolio> = ({ Sectors }) => {
  if (Sectors.length == 0) {
    return null
  }
  return (
    <div className={styles.st_viewport}>
        {Sectors.map((sector, index) => (
          <div className={styles.st_wrap_table} key={index}>
            <header className={styles.st_table_header}>
              <h1 className={styles.h1}>{sector.Sector}</h1>
              <div className={styles.st_row} >
                <div className={styles._name}>Name </div>
                <div className={styles._value}>Value</div>
              </div>
            </header>
            <div className={styles.st_table}>
              <TableRow key={index} sector={sector} />
            </div>
            </div>
        ))}
        </div>
  );
};

const TableRow: React.FC<{ sector: Sector }> = ({ sector }) => {
  return (
    <div className={styles.st_table}>
        {sector.Stocks.map((stock, index) => (
          <div className={styles.st_row} style={{'background':COLORS[sector.Sector]}} key={index}>
            <div className={styles._name}>{stock.Name.slice(0,30)}</div>
            <div className={styles._value}>{formatter.format(stock.Value)}</div>
          </div>
        ))}
    </div>
  );
};

export default Home;