package ru.itis.hr_service.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import ru.itis.hr_service.dto.response.GetCandidatesResponse;
import ru.itis.hr_service.service.HrService;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/ht/sample")
public class HrController {
    private final HrService hrService;

    @GetMapping
    public ResponseEntity<GetCandidatesResponse> getCandidates() {
        return ResponseEntity.ok(GetCandidatesResponse.builder()
                .pairs(hrService.getPairs())
                .build());
    }

}
