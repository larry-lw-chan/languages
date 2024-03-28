interface Mappable {
  location: {
    lat: number;
    lng: number;
  };
  markerContent(): string;
}

class CustomMap {
  mapDiv: HTMLElement;
  private googleMap: google.maps.Map;

  constructor(mapDiv: HTMLElement) {
    this.mapDiv = mapDiv;
    this.googleMap = new google.maps.Map(this.mapDiv, {
      center: { lat: 0, lng: 0 },
      zoom: 1,
    });
  }

  addMarker(mappable: Mappable) {
    const marker = new google.maps.Marker({
      map: this.googleMap,
      position: mappable.location,
    });

    this.addInfoWindow(marker, mappable.markerContent());
    // marker.addListener("click", () => {
    //   console.log(mappable.name);
    //   const infoWindow = new google.maps.InfoWindow({
    //     content: mappable.name,
    //   });
    //   infoWindow.open(this.googleMap, marker);
    // });
  }

  private addInfoWindow(marker: google.maps.Marker, content: string) {
    const infoWindow = new google.maps.InfoWindow({
      content,
    });

    marker.addListener("click", () => {
      infoWindow.open(this.googleMap, marker);
    });
  }
}

export { Mappable, CustomMap };
