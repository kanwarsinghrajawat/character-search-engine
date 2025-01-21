provider "google" {
  project = "rickmorty-search"  
  region  = "us-east1"         
}

resource "google_compute_instance" "mongodb" {
  name         = "rickmorty-mongodb"  
  machine_type = "e2-medium"
  zone         = "us-east1-b"       

  boot_disk {
    initialize_params {
      image = "ubuntu-2004-focal-v20240110"
    }
  }

  attached_disk {
    source      = google_compute_disk.mongodb_disk.id
    device_name = "mongodb-data"
  }

  network_interface {
    network = "default"

    access_config {
      // Allows external access
    }
  }

  metadata_startup_script = <<-EOT
    #!/bin/bash
    sudo apt update
    sudo apt install -y mongodb-org
    sudo systemctl start mongod
    sudo systemctl enable mongod
    # Attach storage
    sudo mkfs.ext4 -F /dev/disk/by-id/google-mongodb-data
    sudo mkdir -p /mnt/mongodb
    sudo mount /dev/disk/by-id/google-mongodb-data /mnt/mongodb
    sudo chown -R mongodb:mongodb /mnt/mongodb
    sudo sed -i 's|/var/lib/mongodb|/mnt/mongodb|' /etc/mongod.conf
    sudo systemctl restart mongod
  EOT

  tags = ["mongodb"]
}

resource "google_compute_disk" "mongodb_disk" {
  name  = "rickmorty-mongodb-disk"  # Updated disk name
  type  = "pd-ssd"
  zone  = "us-east1-b"
  size  = 20
}

resource "google_compute_firewall" "allow_mongodb" {
  name    = "allow-mongodb"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["27017"]
  }

  source_ranges = ["116.75.91.86/32"]  # Replace YOUR_IP with your actual IP address
  target_tags   = ["mongodb"]
}
