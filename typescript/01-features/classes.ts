class Vehicle {
  constructor(public color: string) {}

  drive(): void {
    console.log("chugga chugga");
  }

  protected honk(): void {
    console.log("beep beed");
  }

  getColor(): void {
    console.log(this.color);
  }
}

class Car extends Vehicle {
  constructor(public wheels: number, color: string) {
    super(color);
  }

  startDrivingProcess(): void {
    this.driving();
    this.honk();
  }

  private driving(): void {
    console.log("Vroom");
  }
}

const vehicle = new Vehicle("orange");
vehicle.drive();
vehicle.getColor();

const car = new Car("red");
car.startDrivingProcess();
