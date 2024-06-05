package ru.itis.hr_service.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@NoArgsConstructor
@AllArgsConstructor
@Builder
@Data

public class Cv {
    private String firstName;
    private String lastName;
    private String city;
    private Integer experience;
    private List<String> skills;

}
